package main

import (
	"log"

	"github.com/Johannes-Krabbe/hive-nexus-api/src/common/config"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/common/db"
	"github.com/Johannes-Krabbe/hive-nexus-api/src/common/routes"
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

	h := db.Init(config.DBUrl)

	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{config.ClientUrl}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	// Register the middleware
	router.Use(cors.New(corsConfig))

	routes.RegisterRoutes(router, h)

	// temporary fix for popup in macos
	_ = router.Run("localhost:3001")
}
