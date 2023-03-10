package routes

import (
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/auth"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/comment"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/post"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/post_like"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/pkg/user"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	user.RegisterRoutes(r.Group("/user"), db)
	post.RegisterRoutes(r.Group("/post"), db)
	post_like.RegisterRoutes(r.Group("/like"), db)
	comment.RegisterRoutes(r.Group("/comment"), db)
	auth.RegisterRoutes(r.Group("/auth"), db)
}
