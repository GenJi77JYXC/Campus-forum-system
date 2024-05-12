package routers

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/middleware"
	"Campus-forum-system/model"
	"Campus-forum-system/routers/guest"

	"github.com/gin-gonic/gin"
)

type recoverWriter struct{}

func CollectRouter(r *gin.Engine) *gin.Engine {
	// r.Use(middleware.Cors())
	// r.GET("/ip", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, ctx.ClientIP())
	// })

	// r.POST("/login", controller.Login)
	// r.POST("/regist", controller.Regist)

	// gin.RecoveryWithWriter用于在发生 panic 时恢复程序的执行，并向给定的 io.Writer 写入错误信息。
	r.Use(gin.RecoveryWithWriter(&recoverWriter{}))
	r.Use(middleware.JSONRequestContextHandler(func(c *gin.Context) model.APIRequest {
		if c.Request.URL.Path == "/api/user/register" {
			return new(model.RegisterRequest)
		} else if c.Request.URL.Path == "/api/user/login" {
			return new(model.LoginRequest)
		} else if c.Request.URL.Path == "/api/topics" {
			return new(model.ArticleRequest)
		} else if c.Request.URL.Path == "/api/comments" {
			return new(model.CommentRequest)
		} else if c.Request.URL.Path == "/api/topics/like" {
			return new(model.LikeArticleRequest)
		} else if c.Request.URL.Path == "/api/topics/del_like" {
			return new(model.LikeArticleRequest)
		} else if c.Request.URL.Path == "/api/user/set/username" {
			return new(model.SetUsernameRequest)
		} else if c.Request.URL.Path == "/api/user/set/email" {
			return new(model.SetEmailRequest)
		} else if c.Request.URL.Path == "/api/user/set/password" {
			return new(model.SetPasswordRequest)
		} else if c.Request.URL.Path == "/api/user/update/password" {
			return new(model.UpdatePasswordRequest)
		} else if c.Request.URL.Path == "/api/user/profile" {
			return new(model.UpdateUserProfile)
		} else if c.Request.URL.Path == "/api/topics/favorite" {
			return new(model.FavoriteArticleRequest)
		} else if c.Request.URL.Path == "/api/topics/del_favorite" {
			return new(model.FavoriteArticleRequest)
		}
		return nil
	}))
	r.Use(middleware.ReponseHandler())

	user := r.Group("/api")
	{
		user.GET("/configs", guest.GetConfigs)
		user.POST("/user/register", guest.RegisterByEmail)
		user.POST("/user/login", guest.Login)
		user.GET("/user/logout", guest.Logout)
		user.GET("/user/info/:id", guest.GetUserInfo)
		user.GET("/user/current", guest.GetCurrentUser) // 登录用户信息 5.12 测试到这里（这个地方有问题）发get请求时不论token对不对都会返回true
		user.POST("/user/profile", guest.UpdateUserProfile)
		// user.GET("/user/favorites", guest.GetUserFavorite)
		user.POST("/user/set/username", guest.SetUsername)
		user.POST("/user/set/email", guest.SetEmail)
		user.POST("/user/set/password", guest.SetPassword)
		user.POST("/user/update/password", guest.UpdatePassword)
	}

	return r
}

func (rw *recoverWriter) Write(p []byte) (int, error) {
	logs.Logger.Error(string(p))
	return gin.DefaultErrorWriter.Write(p)
}
