package main

import (
	"log"
	"myapp/config"
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
