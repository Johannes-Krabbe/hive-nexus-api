package user

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
)

type GetUserRequestBody struct {
	Username string `json:"username"`
}

func (h handler) GetUser(c *gin.Context) {
	body := GetUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	// checking if user with username or email already exists
	// using .DB.Limit(1).Find here instead of .First to prevent error messages
	if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "username = ?", body.Username); user.ID <= 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "User with this username does not exist"})
		return
	}

	// TODO discuss what info to share
	c.JSON(http.StatusCreated, gin.H{"data": user})
}
