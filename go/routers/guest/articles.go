package guest

import (
	"Campus-forum-system/logs"
	"Campus-forum-system/model"
	"Campus-forum-system/service"
	"Campus-forum-system/util"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostArticle 发布文章
func PostArticle(c *gin.Context) {
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)

		return
	}

	req := getReqFromContext(c).(*model.ArticleRequest)
	logs.Logger.Info(req)
	var err error
	if req.UserID == 0 {
		err = errors.New("UserID 不存在")
	}
	if !util.CheckContent(req.Content) {
		err = errors.New("内容不能为空")
	}
	if !util.CheckContent(req.Title) {
		err = errors.New("标题不能为空")
	}
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	}
	article, err := service.ArticleService.PostArticle(user, req.Title, req.Content)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, article, "发布成功", true)
}

// GetArticleList 获取文章列表
func GetArticleList(c *gin.Context) {
	logs.Logger.Info(c.Request.URL.Path)
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)
		return
	}
	limit := c.DefaultQuery("limit", "10")
	sortby := c.DefaultQuery("sortby", "created_time")
	order := c.DefaultQuery("order", "desc")
	cursor := c.DefaultQuery("cursor", "2559090472000")
	uID := c.DefaultQuery("user_id", "0")

	var err error
	limitNum, err1 := strconv.Atoi(limit)
	cursorTime, err2 := strconv.ParseInt(cursor, 10, 64)
	authorID, err3 := strconv.ParseInt(uID, 10, 64)
	if err1 != nil || err2 != nil || err3 != nil {
		err = errors.New("参数错误")
	}
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	}
	resp, err := service.ArticleService.GetArticleList(user, authorID, limitNum, cursorTime, sortby, order)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	} else {
		setAPIResponse(c, resp, "获取成功", true)
	}

}

// GetArticleByID 通过ID获取文章详情
func GetArticleByID(c *gin.Context) {
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)
		return
	}
	id := c.Param("id")

	articleID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	}

	resp, err := service.ArticleService.GetArticleByID(user, articleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "获取成功", true)
}
