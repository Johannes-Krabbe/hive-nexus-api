package user

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    r.POST("/create", h.CreateUser)
}