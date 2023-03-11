package router

import (
	"gvb-server/api"
)

// SettingsRouter 注册settings路由
func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi           // 获取settingsApi实例
	router.GET("settings", settingsApi.SettingsInfoView) // 注册路由
}
