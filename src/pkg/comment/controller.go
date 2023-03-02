package comment

import (
	"time"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicCommentData struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
	PostID    uuid.UUID `json:"postID"`
	CommentID uuid.UUID `json:"commentID"`
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
