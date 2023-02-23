package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/helpers"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PublicUserData struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

func (h handler) GetUser(c *gin.Context) {

	// getting params
	username := c.Query("username")
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include Username in querry params")
		return
	}

	var user models.User

	// checking if user with username or email already exists
	// using .DB.Limit(1).Find here instead of .First to prevent error messages
	if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "username = ?", username); user.ID == uuid.Nil {
		fmt.Println(user)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": true, "message": "User with this username does not exist"})
		return
	}

	var viewData PublicUserData
	if err := helpers.TransformData(user, &viewData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Something went wrong"})
		return
	}

	// TODO discuss what info to share
	c.JSON(http.StatusCreated, gin.H{"data": viewData})
}
