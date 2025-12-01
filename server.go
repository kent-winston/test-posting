package main

import (
	"log"
	"myapp/config"
	"myapp/docs"
	"myapp/middleware"
	"myapp/router"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config.ConnectDB()
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	docs.SwaggerInfo.Title = "Posting API"
	docs.SwaggerInfo.Description = "API docs for posting"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()
	r.Use(
		gin.Recovery(),
		middleware.CORSMiddleware(),
		middleware.AuthMiddleware(),
	)

	router.ApiRouter(r)

	log.Println("Listen and serve at http://localhost:" + port)
	r.Run(":" + port)
}
