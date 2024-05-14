package guest

import (
	"Campus-forum-system/model"
	"Campus-forum-system/service"

	"github.com/gin-gonic/gin"
)

func PostLikeArticle(c *gin.Context) {
	user := service.UserService.GetCurrentUser(c)
	if user == nil {
		setAPIResponse(c, nil, "请先登录", false)
		return
	}
	req := getReqFromContext(c).(*model.LikeArticleRequest)
	if req.UserID == 0 || req.ArticleID == 0 {
		setAPIResponse(c, nil, "参数错误", false)
		return
	}

	err := service.LCService.PostLikeArticle(req.UserID, req.ArticleID)
	if err != nil {
		setAPIResponse(c, nil, err.Error(), false)
		return
	}
	setAPIResponse(c, nil, "操作成功", true)

}
