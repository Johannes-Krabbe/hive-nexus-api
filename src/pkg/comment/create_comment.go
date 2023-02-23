package comment

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
)

type CreateCommentRequestBody struct {
	Content string `json:"Content"`
	PostID  uint   `json:"PostID"`
}

func (h handler) CreateComment(c *gin.Context) {
	body := CreateCommentRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userID, ok := c.Get("UserID")
	if !ok { // why internal server error?
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is missing from context"})
	}

	var post = models.Post
	var user = models.User

	if result := h.DB.Find(&user, userID); result.Error != nil || user.ID <= 0 {
		c.AbortWithError(http.StatusUnauthorized)
		return
	}

	var postID = body.PostID
	if result := h.DB.Find(&post, postID); result.Error != nil || post.ID <= 0 {
		c.AbortWithError(http.StatusBadRequest)
		return
	}

	var comment = models.Comment
	comment.Content = body.Content
	comment.User = user
	comment.Post = post

	if result := h.DB.Create(&comment); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &comment)
}
