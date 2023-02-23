package post

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPosts(c *gin.Context) {

	var posts []models.Post

	h.DB.Preload("User").Find(&posts)

	c.JSON(http.StatusCreated, gin.H{"data": posts})
}
