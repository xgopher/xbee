package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xgopher/xbee/modules/common"
	_ "github.com/xgopher/xbee/routers"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 静态文件
	beego.SetStaticPath("/static", "web")

	beego.Run()
}
