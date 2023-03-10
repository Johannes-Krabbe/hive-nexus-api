package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublicUserData struct {
	Username  string    `json:"username,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.GET("/one", h.GetUser)
}
