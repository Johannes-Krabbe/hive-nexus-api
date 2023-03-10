package comment

import (
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) GetComments(c *gin.Context) {

	// getting params
	postID := c.Query("postID")
	userID := c.Query("userID")
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if postID == "" && userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include postID or userID in query params")
		return
	}

	if userID != "" && postID != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Include just one: Username or Email in query params")
		return
	}

	if userID != "" {
		var user models.User
		var comments []models.Comment

		if h.DB.Limit(1).Select("ID", "CreatedAt", "Username").Find(&user, "ID = ?", userID); user.ID == uuid.Nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, "User with this userID does not exist.")
			return
		}
		h.DB.Find(&comments, "user_id = ?", user.ID)
		c.JSON(http.StatusOK, gin.H{"data": comments})
		return

	}

	var post models.Post
	if h.DB.Limit(1).Find(&post, "ID = ?", postID); post.ID == uuid.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Post with this postID does not exist.")
		return
	}

	var comments []models.Comment
	h.DB.Joins(
		"User", h.DB.Select(
			"ID", "username",
		),
	).Find(
		&comments, "post_id = ?", postID,
	)

	var viewDataArr []PublicCommentData

	for _, comment := range comments {
		var viewData PublicCommentData

		viewData.CommentID = comment.ID
		viewData.PostID = comment.PostID
		viewData.Content = comment.Content
		viewData.CreatedAt = comment.CreatedAt
		viewData.Username = comment.User.Username

		viewDataArr = append(viewDataArr, viewData)
	}

	c.JSON(http.StatusOK, gin.H{"data": viewDataArr})

}
