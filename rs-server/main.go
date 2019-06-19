package main

import (
	"github.com/astaxie/beego"
	"rs-engine/rs-server/controllers"
)

func main() {
	beego.Router("/", &controllers.ApiController{}, "get:Version")
	beego.Run()
}