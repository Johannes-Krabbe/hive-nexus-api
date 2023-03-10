package post

import (
	"time"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type PublicPostData struct {
	PostID    uuid.UUID `json:"postID,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Username  string    `json:"username,omitempty"`
}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.Use(auth.VerifyTokenMiddleware())

	r.POST("/create", h.CreatePost)
	r.POST("/delete", h.DeletePost)
	r.GET("/one", h.GetPost)
	r.GET("/all", h.GetPosts)
}
