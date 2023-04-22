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

	var viewDataArr []PublicPostData

	for _, post := range posts {
		var viewData PublicPostData

		viewData.PostID = post.ID
		viewData.Title = post.Title
		viewData.Content = post.Content
		viewData.CreatedAt = post.CreatedAt
		viewData.Username = post.User.Username

		viewDataArr = append(viewDataArr, viewData)
	}

	c.JSON(http.StatusOK, gin.H{"data": viewDataArr})
}
