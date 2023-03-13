package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type PublicUserData struct {
	ID        uuid.UUID `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Token     string    `json:"token,omitempty"`
}

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r.POST("/sign-up", h.SignUp)
	r.POST("/sign-in", h.SignIn)
	r.GET("/availability-check", h.AvailabilityCheck)
}
