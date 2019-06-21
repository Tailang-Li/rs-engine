package routers

import (
	"github.com/astaxie/beego"
	"rs-engine/rs-server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/api",
			beego.NSRouter("/version", &controllers.ApiController{}, "get:Version"),
		),
	)
	beego.AddNamespace(ns)
}