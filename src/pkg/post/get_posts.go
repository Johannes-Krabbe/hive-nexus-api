package post

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetPosts(c *gin.Context) {
	var posts []models.Post

	if result := h.DB.Preload("User").Find(&posts); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var viewData []PublicPostData

	for _, post := range posts {
		var p PublicPostData

		p.Content = post.Content
		p.Title = post.Title
		p.CreatedAt = post.CreatedAt
		p.Username = post.User.Username

		viewData = append(viewData, p)
	}

	c.JSON(http.StatusCreated, gin.H{"data": viewData})
}
