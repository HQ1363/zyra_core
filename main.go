package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	_ "zyra_core/routers"
	"zyra_core/tasks"
	"zyra_core/utils"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"zyra_core/models"
)

var logger = logs.GetLogger("main")

func init() {
	models.InitDB()
	models.InitConn()
	utils.InitLogs()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else if beego.BConfig.RunMode == "prod" {
		logger.Println("current env in prod.")
		// in prod we need to what
		testCronTask := toolbox.NewTask("testTask", "*/2 * * * * *", func() error {
			tasks.TestTask()
			return nil
		})
		toolbox.AddTask("testTask", testCronTask)
		defer toolbox.StopTask()
	}
	toolbox.StartTask()
	beego.Run()
}
