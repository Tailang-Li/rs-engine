package conf

import (
	"database/sql"
	"github.com/astaxie/beego"
)

func init() {
	url := beego.AppConfig.String("mysql::url")

	db, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
}
