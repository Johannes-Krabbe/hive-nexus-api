package comment

import (
	"errors"
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DeleteCommentRequestBody struct {
	CommentID uuid.UUID `json:"commentID"`
}

func (h handler) DeleteComment(c *gin.Context) {
	body := DeleteCommentRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, ok := c.Get("UserID")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is missing from context"})
		return
	}

	var comment models.Comment
	if result := h.DB.Find(&comment, body.CommentID); result.Error != nil || comment.ID == uuid.Nil {
		c.AbortWithError(http.StatusUnauthorized, result.Error)
		return
	}

	if comment.UserID != userID {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized request to delete comment."))
		return
	}

	if result := h.DB.Delete(&comment); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var viewData PublicCommentData
	viewData.CommentID = comment.ID
	viewData.Content = comment.Content
	viewData.CreatedAt = comment.CreatedAt
	viewData.Username = comment.User.Username
	viewData.PostID = comment.PostID

	c.JSON(http.StatusOK, gin.H{"data": viewData})
}
