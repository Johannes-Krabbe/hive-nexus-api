package comment

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateCommentRequestBody struct {
	Content string    `json:"content"`
	PostID  uuid.UUID `json:"postID"`
}

func (h handler) CreateComment(c *gin.Context) {
	body := CreateCommentRequestBody{}

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
	var user models.User

	if result := h.DB.Find(&user, userID); result.Error != nil || user.ID == uuid.Nil {
		c.AbortWithError(http.StatusUnauthorized, result.Error)
		return
	}

	var postID = body.PostID
	if result := h.DB.Find(&post, postID); result.Error != nil || post.ID == uuid.Nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	var comment models.Comment
	comment.Content = body.Content
	comment.User = user
	comment.Post = post

	if result := h.DB.Create(&comment); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	var viewData PublicCommentData
	viewData.Content = comment.Content
	viewData.CreatedAt = comment.CreatedAt
	viewData.Username = comment.User.Username
	viewData.PostID = comment.PostID

	c.JSON(http.StatusCreated, gin.H{"data": viewData})
}
