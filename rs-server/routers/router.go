package routers

import (
	"rs-engine/rs-server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/api",
			beego.NSRouter("/version", &controllers.ApiController{}, "get:Version"),
		),
		beego.NSNamespace("/test",
			beego.NSRouter("/save", &controllers.TestModelController{}, "get:Save"),
			beego.NSRouter("/update", &controllers.TestModelController{}, "get:Update"),
			beego.NSRouter("/read", &controllers.TestModelController{}, "get:Read"),
			beego.NSRouter("/delete", &controllers.TestModelController{}, "get:Delete"),
		),
	)
	beego.AddNamespace(ns)
}
