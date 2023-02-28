package post_like

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePostLikeRequestBoday struct {
	PostID uuid.UUID `json:"PostID"`
}

func (h handler) CreatePostLike(c *gin.Context) {
	body := CreatePostLikeRequestBoday{}

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
	var postLike models.PostLike

	if result := h.DB.Find(&post, body.PostID); result.Error != nil || post.ID == uuid.Nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	if result := h.DB.Find(&user, userID); result.Error != nil || user.ID == uuid.Nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	postLike.User = user
	postLike.Post = post

	if result := h.DB.Create(&postLike); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &postLike)
}
