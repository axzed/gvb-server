package router

import (
	"github.com/gin-gonic/gin"
	"gvb-server/global"
)

// enter 为入口文件，用于初始化总路由

type RouterGroup struct {
	*gin.RouterGroup
}

// InitRouter 初始化总路由
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) // 设置gin运行模式
	router := gin.Default()               // 初始化默认gin路由

	apiRouterGroup := router.Group("/api") // 初始化api路由组

	routerGroupApp := RouterGroup{apiRouterGroup} // 初始化路由组
	// 系统配置api
	routerGroupApp.SettingsRouter() // 注册settings路由

	return router
}
