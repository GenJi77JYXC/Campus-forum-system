package routers

import (
	"Campus-forum-system/controller"
	"Campus-forum-system/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.Use(middleware.Cors())
	r.GET("/ip", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, ctx.ClientIP())
	})

	r.POST("/login", controller.Login)
	r.POST("/regist", controller.Regist)

	return r
}
