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
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include PostID in query params")
		return
	}

	var post models.Post

	if postID != "" {
		if h.DB.Limit(1).Select("ID", "CreatedAt", "Title", "Content").Find(&post, "ID = ?", postID); post.ID == uuid.Nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": true, "message": "Post with this PostID does not exist"})
			return
		}
	}

	var viewData PublicPostData

	viewData.Username = post.User.Username
	viewData.Title = post.Title
	viewData.Content = post.Content
	viewData.CreatedAt = post.CreatedAt
	viewData.CreatedAt = post.CreatedAt

	c.JSON(http.StatusOK, gin.H{"data": viewData})
}
