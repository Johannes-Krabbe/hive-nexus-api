package auth

import (
	"fmt"
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignInRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h handler) SignIn(c *gin.Context) {
	body := SignInRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	// checking if user with username or email already exists
	// using .DB.Limit(1).Find here instead of .First to prevent error messages
	if h.DB.Limit(1).Find(&user, "email = ?", body.Email); user.ID == uuid.Nil {
		fmt.Println(user)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Email does not exists"})
		return
	}

	inPassHash, err := hash(body.Password, user.Salt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Something went wrong"})
		return
	}

	if inPassHash != user.Password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Password is wrong"})
		return
	}

	token, err := generateJWT(user.ID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Something went wrong"})
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
