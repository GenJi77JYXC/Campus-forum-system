package guest

import (
	"Campus-forum-system/model"
	"Campus-forum-system/service"
	"Campus-forum-system/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)
		return
	}
	req := getReqFromContext(c).(*model.CommentRequest)
	req.Content = util.DeletePreAndSufSpace(req.Content)
	if req.UserID == 0 || req.ArticleID == 0 || req.Content == "" {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}

	resp, err := service.CommentService.BuildComment(user.ID, req.ArticleID, req.ParentID, req.Content)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
	}
	setAPIResponse(c, resp, "评论成功", true)
}

func GetComments(c *gin.Context) {
	id := c.Query("article_id")
	cursor := c.DefaultQuery("cursor", "2559090472000")
	cursorTime, err1 := strconv.ParseInt(cursor, 10, 64)
	articleID, err2 := strconv.ParseInt(id, 10, 64)
	if err1 != nil || err2 != nil {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}
	resp, err := service.CommentService.GetCommentList(articleID, cursorTime)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "查询成功", true)
}

func GetLikeComments(c *gin.Context) {
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)
		return
	}
	id := c.Query("comments_id")
	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}
	resp, err := service.CommentService.LikeComment(commentID, user.ID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "点赞成功", true)
}

func CancelCommentLike(c *gin.Context) {
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)
		return
	}
	id := c.Query("comments_id")
	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}
	resp, err := service.CommentService.UnlikeComment(commentID, user.ID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, resp, "取消点赞成功", true)
}

func DeleteCommentByID(c *gin.Context) {
	user, err1 := service.UserService.GetCurrentUser(c)
	if user == nil || err1 != nil {
		setAPIResponse(c, nil, err1.Error(), false)
		return
	}
	id := c.Query("comment_id")
	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}
	err = service.CommentService.DeleteComment(commentID, user.ID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, nil, "删除成功", true)
}
