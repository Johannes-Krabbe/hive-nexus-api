package auth

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SignUpRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var validate *validator.Validate

func (h handler) SignUp(c *gin.Context) {
	validate = validator.New()
	body := SignUpRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}

	var user models.User

	//checking if user with username or email already exists
	// using .DB.Limit(1).Find here instead of .First to prevent error messages
	if h.DB.Limit(1).Find(&user, "username = ?", body.Username); user.ID != uuid.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Username already exists"})
		return
	}

	if h.DB.Limit(1).Find(&user, "email = ?", body.Email); user.ID != uuid.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Email already exists"})
		return
	}

	// creating the user
	user.Username = body.Username
	user.Email = body.Email

	user.Salt = getSalt()

	pw, err := hash(body.Password, user.Salt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Something went wrong"})
		return
	}

	user.Password = pw

	// validate user data
	err = validate.Struct(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err, "message": "Validation Error"})
		return
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err, "message": "Validation Error"})
		return
	}

	var viewData PublicUserData
	viewData.ID = user.ID
	viewData.CreatedAt = user.CreatedAt
	viewData.Username = user.Username
	viewData.Email = user.Email
	viewData.Token = token

	c.JSON(http.StatusCreated, gin.H{"data": viewData})
}
