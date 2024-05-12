package guest

import (
	"Campus-forum-system/model"

	"github.com/gin-gonic/gin"
)

// GetConfigs return config of server
func GetConfigs(c *gin.Context) {
	resp := model.SysConfigResponse{}
	resp.SiteTitle = "校园论坛"
	resp.SiteDescription = "学生交流平台"
	resp.SiteNavs = []model.ActionLink{{
		Title: "技术交流",
		URL:   "http://localhost:3000/",
	}, {
		Title: "社会百态",
		URL:   "http://localhost:3000/",
	}}
	resp.TokenExpireDays = 2
	resp.SiteKeywords = []string{"交流", "分享"}
	setAPIResponse(c, resp, "获取配置成功", true)
}
