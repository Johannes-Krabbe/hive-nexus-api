package user

import (
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
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include Username in querry params")
		return
	}

	var user models.User

	if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "username = ?", username); user.ID == uuid.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "User with this username does not exist"})
		return
	}

	var viewData PublicUserData
	if err := helpers.TransformData(user, &viewData); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": true, "message": "Something went wrong"})
		return
	}

	// TODO discuss what info to share
	c.JSON(http.StatusOK, gin.H{"data": viewData})
}
