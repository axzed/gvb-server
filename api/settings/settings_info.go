package settings

import (
	"github.com/gin-gonic/gin"
	"gvb-server/global"
	"gvb-server/model/res"
)

// SettingsInfoView 查寻网站信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
