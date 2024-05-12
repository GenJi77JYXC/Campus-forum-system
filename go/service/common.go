package service

import (
	"Campus-forum-system/model"

	"github.com/gin-gonic/gin"
)

func getReqFromContext(c *gin.Context) interface{} {
	req, _ := c.Get(model.CTXAPIReq)
	return req
}
