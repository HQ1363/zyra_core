package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 日志记录器
var Logger = logs.GetLogger("utils")

// consoleLogs开发模式下日志
var consoleLogs *logs.BeeLogger

// fileLogs 生产环境下日志
var fileLogs *logs.BeeLogger

func InitLogs() {
	Logger.Println("start setting log output")
	if beego.BConfig.RunMode == "dev" {
		consoleLogs = logs.NewLogger(1)
		err := consoleLogs.SetLogger(logs.AdapterConsole, `{"level": 1,"color": true}`)
		if err != nil {
			Logger.Println("设置日志输出至终端失败: ", err)
		}
		consoleLogs.Async()  //异步
	} else if beego.BConfig.RunMode == "prod" {
		fileLogs = logs.NewLogger(10000)
		level := beego.AppConfig.String("logs::level")
		err := fileLogs.SetLogger(logs.AdapterMultiFile, `{
		"filename":"logs/zyra_core.log",
		"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"],
		"level":`+level+`,
		"daily":true,
		"maxdays":10}`)
		if err != nil {
			Logger.Println("设置日志输出至文件失败: ", err)
		}
		fileLogs.Async()  //异步
	}
}

//Log 输出日志
func log(level, v interface{}) {
	format := "%s"
	if level == "" {
		level = "debug"
	}
	switch level {
		case "emergency":
			fileLogs.Emergency(format, v)
		case "alert":
			fileLogs.Alert(format, v)
		case "critical":
			fileLogs.Critical(format, v)
		case "error":
			fileLogs.Error(format, v)
		case "warning":
			fileLogs.Warning(format, v)
		case "notice":
			fileLogs.Notice(format, v)
		case "info":
			fileLogs.Info(format, v)
		case "debug":
			fileLogs.Debug(format, v)
		case "trace":
			fileLogs.Trace(format, v)
		default:
			fileLogs.Debug(format, v)
	}
	switch level {
		case "emergency":
			consoleLogs.Emergency(format, v)
		case "alert":
			consoleLogs.Alert(format, v)
		case "critical":
			consoleLogs.Critical(format, v)
		case "error":
			consoleLogs.Error(format, v)
		case "warning":
			consoleLogs.Warning(format, v)
		case "notice":
			consoleLogs.Notice(format, v)
		case "info":
			consoleLogs.Info(format, v)
		case "debug":
			consoleLogs.Debug(format, v)
		case "trace":
			consoleLogs.Trace(format, v)
		default:
			consoleLogs.Debug(format, v)
	}
}
