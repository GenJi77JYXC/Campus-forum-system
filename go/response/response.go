package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, code int, data, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Fail(ctx *gin.Context, code int, data, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Response(ctx *gin.Context, httpstatus, code int, data, msg any) {
	ctx.JSON(httpstatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
