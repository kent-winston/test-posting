package main

import (
	"log"
	"myapp/config"
	"myapp/middleware"
	"myapp/router"
	"os"

	_ "myapp/docs" // Import docs for Swagger

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const defaultPort = "8080"

// @title 		Test Posting API
// @version 	1.0
// @description This is a test API for Test Posting
// @host 		localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config.ConnectDB()
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.CORSMiddleware(),
		middleware.AuthMiddleware(),
	)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.ApiRouter(r)

	log.Println("Listen and serve at http://localhost:" + port)
	r.Run(":" + port)
}
