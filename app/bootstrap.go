package app

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func Run()  {
	loadConfig()
	loadMiddleware()
	logs.Info(beego.AppConfig.String("dbconn"))
	beego.Run()
}

func loadConfig()  {
	beego.LoadAppConfig("ini", "config/app.ini")
}

// 加载中间件，需要自定义的中间可以在这个函数添加
func loadMiddleware()  {
	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		logs.Info("111")
		//json := map[string]string{"hello":"dd"}
		//context.Output.JSON(json, true, true)
	})

	beego.InsertFilter("*", beego.BeforeRouter, func(context *context.Context) {
		logs.Info("222")
		json := map[string]string{"hello":"dd"}
		context.Output.JSON(json, true, true)
	})
}
