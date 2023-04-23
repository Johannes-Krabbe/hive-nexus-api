package main

import (
	"log"
	// "net/http"

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

	log.Println("Configure CORS")
	corsConfig := cors.DefaultConfig()

	corsConfig.AllowOrigins = []string{config.ClientUrl}
	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true

	// OPTIONS method for ReactJS
	// corsConfig.AddAllowMethods("OPTIONS")

	corsConfig.AddAllowHeaders("Authorization")
	log.Println("AllowHeadders: ", corsConfig.AllowHeaders)
	log.Println("AllowOrigins: ", corsConfig.AllowOrigins)
	// Register the middleware
	router.Use(cors.New(corsConfig))
	log.Println("Success: Configure CORS")

	routes.RegisterRoutes(router, h)

	/*
		var test = "clientUrl: " + config.ClientUrl
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, test)
		})
	*/

	// temporary fix for popup in macos
	_ = router.Run()
}
