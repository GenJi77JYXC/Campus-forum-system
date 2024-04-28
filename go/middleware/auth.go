package middleware

import (
	"Campus-forum-system/database"
	"Campus-forum-system/model"
	"Campus-forum-system/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求头拿到Authorization
		tokenString := ctx.GetHeader("Authorization")
		// fmt.Println("authMiddleWare中的token:" + tokenString)
		// 看前缀 Bearer xxxx
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			response.Response(ctx, http.StatusUnauthorized, 401, "token失效", "身份验证失效")
			ctx.Abort()
			return
		}

		// 剪掉前7个 Bearer和一个空格
		tokenString = tokenString[7:]

		// 在redis黑名单中查看token是否已经注销过
		// if database.RedisGetKey(tokenString) != "" {
		// 	// 这个token在redis中， 表示已经注销过了
		// 	response.Response(ctx, http.StatusUnauthorized, 401, "token被注销", "身份失效，重新登陆")
		// 	ctx.Abort()
		// 	return
		// }

		//解token
		token, claims, err := ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.Response(ctx, 500, 500, "token解析错误", "用户信息失效")
			fmt.Println("token解析错误：", err)
			ctx.Abort()
			return
		}
		//从claims中拿信息
		userID := claims.UserID
		// 拿数据库数据
		var user model.User
		DB := database.GetDB()
		DB.Table("users").Where("id = ?", userID).First(&user)
		if user.ID == 0 {
			response.Response(ctx, 401, 401, "用户不存在", "用户不存在")
			ctx.Abort()
			return
		}
		// 把东西放到上下文
		ctx.Set("user", user)
		ctx.Set("token", tokenString)
		ctx.Next()
	}
}
