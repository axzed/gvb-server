package settings

import (
	"github.com/gin-gonic/gin"
	"gvb-server/model/res"
)

// SettingsInfoView 是settings的一个具体路由对应的方法
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.Ok("成功", map[string]any{
		"system_name": "GVB",
	}, c)
}
