package post_like

import (
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicPostLikeData struct {
	PostLikeID uuid.UUID `json:"postLikeID"`
	UserID     uuid.UUID `json:"userID"`
	PostID     uuid.UUID `json:"postID"`
}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.Use(auth.VerifyTokenMiddleware())

	r.POST("/create", h.CreatePostLike)
	r.POST("/delete", h.DeletePostLike)
}
