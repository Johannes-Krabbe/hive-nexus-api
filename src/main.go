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
    h := db.Init(config.DBUrl)


    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{config.ClientUrl}
    corsConfig.AllowCredentials = true
  
    router.Use(cors.New(corsConfig))


    routes.RegisterRoutes(router, h)

    router.Run(config.Port)
}