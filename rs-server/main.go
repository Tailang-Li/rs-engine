package main

import (
	"github.com/astaxie/beego"
	_ "rs-engine/rs-server/routers"
)

func main() {
	beego.Run()
}