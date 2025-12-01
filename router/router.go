package router

import (
	"myapp/controller"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApiRouter(r *gin.Engine) {
	r.POST("/auth/refresh", controller.AuthRefreshToken)

	r.POST("/post", controller.PostCreate)
	r.PUT("/post", controller.PostUpdate)
	r.DELETE("/post", controller.PostDelete)
	r.GET("/posts", controller.PostGetAll)
	r.GET("/post", controller.PostGetByID)

	r.POST("/user/register", controller.UserRegister)
	r.POST("/user/login", controller.UserLogin)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoute := r.Group("")
	authRoute.Use(middleware.IsLogin())
	authRoute.GET("/user/me", controller.UserGetMe)
}
