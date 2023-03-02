package post_like

import (
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PublicPostLikeData struct {
	PostLikeID uuid.UUID `json:"postLikeID,omitempty"`
	UserID     uuid.UUID `json:"userID,omitempty"`
	PostID     uuid.UUID `json:"postID,omitempty"`
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
