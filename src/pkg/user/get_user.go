package user

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) GetUser(c *gin.Context) {

	// getting params
	username := c.Query("username")
	userID := c.Query("userID")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if username == "" && userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include username or userID in query params")
		return
	}

	if username != "" && userID != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include just one: username or userID in query params")
		return
	}

	var user models.User
	if username != "" {
		if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "username = ?", username); user.ID == uuid.Nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "User with this username does not exist"})
			return
		}
	} else {
		if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "ID = ?", userID); user.ID == uuid.Nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "User with this ID does not exist"})
			return
		}
	}

	var viewData PublicUserData

	viewData.Username = user.Username
	viewData.CreatedAt = user.CreatedAt

	// TODO discuss what info to share
	c.JSON(http.StatusOK, gin.H{"data": viewData})
}
