package router

import (
	"myapp/controller"
	"myapp/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r.POST("/post/create", controller.PostCreate)
	r.PUT("/post/update", controller.PostUpdate)
	r.DELETE("/post/delete", controller.PostDelete)
	r.GET("/posts", controller.PostGetAll)
	r.GET("/post", controller.PostGetByID)

	r.POST("/user/register", controller.UserRegister)
	r.POST("/user/login", controller.UserLogin)

	authRoute := r.Group("")
	authRoute.Use(middleware.IsLogin())
	authRoute.GET("/user/me", controller.UserGetMe)
}
