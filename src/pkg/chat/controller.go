package chat

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublicChatData struct{}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	// r.Use(auth.VerifyTokenMiddleware())

	r.GET("/start", h.StartChat)
}
