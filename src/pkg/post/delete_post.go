package post

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeletePostRequestBody struct {
	PostID uuid.UUID `json:"postID"`
}

func (h handler) DeletePost(c *gin.Context) {
	body := DeletePostRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, ok := c.Get("UserID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is missing from context"})
		return
	}

	var post models.Post

	if result := h.DB.Find(&post, body.PostID); result.Error != nil || post.ID == uuid.Nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if post.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authorized to delete this post"})
		return
	}

	if result := h.DB.Delete(&post); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var viewData PublicPostData
	viewData.PostID = post.ID

	c.JSON(http.StatusCreated, gin.H{"data": viewData})
}
