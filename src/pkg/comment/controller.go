package comment

import (
	"time"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicCommentData struct {
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Username  string    `json:"username,omitempty"`
	PostID    uuid.UUID `json:"postID,omitempty"`
	CommentID uuid.UUID `json:"commentID,omitempty"`
}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.Use(auth.VerifyTokenMiddleware())

	r.POST("/create", h.CreateComment)
	r.DELETE("/delete", h.DeleteComment)
	r.GET("/get-multiple", h.GetComments)
}
