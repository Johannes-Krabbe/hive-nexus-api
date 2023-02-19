package user

import (
	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type AddUserRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var validate *validator.Validate

func (h handler) CreateUser(c *gin.Context) {
	validate = validator.New()
	body := AddUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Username = body.Username
	user.Email = body.Email
	user.Password = body.Password

	err := validate.Struct(user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}
