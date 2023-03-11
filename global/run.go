package global

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Run 优雅启停函数
// r: gin.Engine实例, srvName: 服务名称, addr: 监听地址
func Run(r *gin.Engine, srvName string, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	// 开启goroutine启动服务 保证主线程不阻塞
	go func() {
		Log.Infof("%s running in %s \n", srvName, srv.Addr)
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			Log.Fatalln(err)
		}
	}()

	// 监听退出信号
	quit := make(chan os.Signal)
	//SIGINT 用户发送INTR字符(Ctrl+C)触发 kill -2
	//SIGTERM 结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // signal.Notify()函数用于监听信号
	<-quit                                               // 阻塞直到有信号传入
	Log.Infof("Shutting Down project %s...", srvName)

	// 优雅关闭服务
	// 2秒内处理完请求
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 执行自定义函数
	if stop != nil {
		stop()
	}
	// 关闭服务
	if err := srv.Shutdown(ctx); err != nil {
		Log.Fatalf("%s Shutdown, caused by: %v", srvName, err)
	}
	// 等待2秒
	select {
	case <-ctx.Done(): // ctx.Done()会在ctx.Done()被关闭或者超时的时候返回一个空struct{}
		Log.Infof("wait time out...")
	}
	Log.Infof("%s exiting...", srvName)
}
