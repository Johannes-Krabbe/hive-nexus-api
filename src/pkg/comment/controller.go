package comment

import (
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
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

	r.Use(auth.VerifyTokenMiddleware())

	r.POST("/create", h.CreateComment)
	r.GET("/get-multiple", h.GetComments)
	// r.DELETE("/delete", h.DeleteComment)
}
