package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb-server/global"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	// 如果entry.Buffer不为空，就直接使用，否则就新建一个
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	log := global.Config.Logger // 从全局变量中获取配置

	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// HasCaller()判断是否开启了返回函数名和行号
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line) // 获取文件名和行号
		//自定义输出格式
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	// 返回 []byte
	return b.Bytes(), nil
}

// InitLogger is the function to initialize the logger.
func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                //新建一个实例
	mLog.SetOutput(os.Stdout)                           //设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                  //设置自己定义的Formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level) //设置最低的Level
	InitDefaultLogger()  // InitDefaultLogger()初始化全局log
	return mLog
}

// InitDefaultLogger is the function to initialize the default logger.
func InitDefaultLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                  //设置自己定义的Formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的Level
}
