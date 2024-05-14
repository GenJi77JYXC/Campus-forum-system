package middleware

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAPIRequestModel returns the model used to store request info
type GetAPIRequestModel func(*gin.Context) model.APIRequest

// JSONRequestContextHandler is the middleware to preproccess request
func JSONRequestContextHandler(getAPIRequestModel GetAPIRequestModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqMethod := c.Request.Method

		if reqMethod == "GET" {
			params := c.Request.URL.Query()
			c.Set(model.CTXAPIURLParams, params)
		} else if reqMethod == "POST" {
			requestBody, err := io.ReadAll(c.Request.Body) //func ReadAll(r Reader) ([]byte, error)  ReadAll 从 r 读取，直到出现错误或 EOF，并返回它读取的数据。 成功的调用返回 err == nil，而不是 err == EOF。因为 ReadAll 是 定义为从 src 读取到 EOF，它不会处理从读取的 EOF 作为要报告的错误。
			if err != nil {
				c.Abort()
			} else {
				c.Set(model.CTXCacheBody, requestBody)
			}
			c.Request.Body = io.NopCloser(bytes.NewReader(requestBody)) // NopCloser 返回一个带有 no-op Close 方法包装的 ReadCloser 提供的读取器 r。 如果 r 实现 WriterTo，则返回的 ReadCloser 将通过转发对 r 的调用来实现 WriterTo。

			params := c.Request.URL.Query()
			c.Set(model.CTXAPIURLParams, params)
			c.Set(model.CTXAPICacheBody, requestBody)
			req := getAPIRequestModel(c)
			if req != nil {
				// parse json to struct 通过BindJSON()可见将json请求体绑定到一个结构体上。
				// 通过 BindJSON () 可见将 json 请求体绑定到一个结构体上。。
				if err = c.BindJSON(req); err == nil {
					c.Set(model.CTXAPIReq, req)
				} else {
					logs.Logger.Error("parse json error:", err)
					c.Abort()
				}
			} else {
				logs.Logger.Error("program can not find the struct matched the request!")
				c.Abort()
			}
		}

		c.Next()
	}
}

// ReponseHandler is the middleware to fill response at the end of program execution
func ReponseHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()

		logs.Logger.Info("The time cost is ", end.Sub(start).Nanoseconds()/1000000)

		var resp *model.APIResponse = new(model.APIResponse)
		if c.IsAborted() {
			resp.Code = 500
			resp.Message = "Program runtime error."
			resp.Value = nil
		} else {
			resp.Code = 1000
			if message, exist := c.Get(model.CTXAPIResponseMessage); exist {
				resp.Message = message.(string)

			}
			if value, exist := c.Get(model.CTXAPIResponseValue); exist {
				resp.Value = value
			}
			if success, exist := c.Get(model.CTXAPIResponseSuccess); exist {
				resp.Success = success.(bool)
			}
		}
		if resp.Code == 500 {
			c.AbortWithStatusJSON(400, resp)
		}
		c.JSON(200, resp)
	}
}
