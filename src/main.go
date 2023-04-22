package main

import (
	"log"
	"net/http"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/common/config"
	// "github.com/Johannes-Krabbe/hive-nexus-api/src/common/db"
	// "github.com/Johannes-Krabbe/hive-nexus-api/src/common/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	// Disable automatic redirects
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	// h := db.Init(config.DBUrl)
	// h := db.Init("postgres://postgres:postgres@172.22.0.2:5432/hive-nexus-api")

	corsConfig := cors.DefaultConfig()

	// corsConfig.AllowOrigins = []string{config.ClientUrl}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	// router.Use(cors.New(corsConfig))

	// routes.RegisterRoutes(router, h)

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, config.ClientUrl)
	})

	// temporary fix for popup in macos
	_ = router.Run()
}
