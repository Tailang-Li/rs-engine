package main

import (
	_ "rs-engine/rs-server/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "rs-engine/rs-server/conf"
)

func main() {
	beego.Run()
}
