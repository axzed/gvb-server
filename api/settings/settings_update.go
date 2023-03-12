package settings

import (
	"github.com/gin-gonic/gin"
	"gvb-server/config"
	"gvb-server/core"
	"gvb-server/global"
	"gvb-server/model/res"
)

// SettingsInfoUpdateView 更新网站信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr config.SiteInfo
	// 从请求中获取json数据并绑定到cr
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ErrorArgument, c)
		return
	}
	global.Config.SiteInfo = cr // 将绑定的修改好的系统配置信息赋值给全局变量
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
	}
	res.OkWith(c)
}
