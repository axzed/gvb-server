package main

import (
	"gvb-server/core"
	"gvb-server/flag"
	"gvb-server/global"
	"gvb-server/router"
)

func main() {
	core.InitConf()                // 初始化配置文件
	global.Log = core.InitLogger() // 初始化日志
	global.DB = core.InitGorm()    // 初始化gorm
	// 命令行参数绑定
	option := flag.Parse()      // 解析命令行参数
	if flag.IsWebStop(option) { // 判断是否停止web项目
		flag.SwitchOption(option) // 根据命令执行不同的函数
		return
	}
	router := router.InitRouter()                // 初始化路由
	srvName := global.Config.System.GetSrvName() // 获取服务名称
	addr := global.Config.System.Addr()          // 获取启动地址
	global.Run(router, srvName, addr, nil)       // 启动服务 (调用global中的Run函数 优雅启停)
}
