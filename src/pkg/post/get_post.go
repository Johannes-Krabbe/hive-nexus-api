package post

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) GetPost(c *gin.Context) {
	// getting params
	postID := c.Query("postId")

	if postID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include postId in query params")
		return
	}

	var post models.Post
	if postID != "" {
		if h.DB.Preload("User").Limit(1).Select("ID", "Title", "Content", "CreatedAt", "UserID").Find(&post, "ID = ?", postID); post.ID == uuid.Nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Post with this PostID does not exist"})
			return
		}
	}

	var viewData PublicPostData

	viewData.PostID = post.ID
	viewData.Title = post.Title
	viewData.Content = post.Content
	viewData.CreatedAt = post.CreatedAt
	viewData.Username = post.User.Username

	c.JSON(http.StatusOK, gin.H{"data": viewData})
}
