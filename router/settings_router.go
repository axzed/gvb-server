package router

import (
	"gvb-server/api"
)

// SettingsRouter 注册settings路由
func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi                  // 获取settingsApi实例
	router.GET("/settings", settingsApi.SettingsInfoView)       // 查询网站信息
	router.PUT("/settings", settingsApi.SettingsInfoUpdateView) // 更新网站信息
}
