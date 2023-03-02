package post

import (
	"time"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type PublicPostData struct {
	PostID    uuid.UUID `json:"postID"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `json:"username"`
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
	r.GET("/all", h.GetPosts)
}
