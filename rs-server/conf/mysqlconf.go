package conf

import (
	"github.com/astaxie/beego/orm"
	"rs-engine/rs-server/models"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local", 30)
	orm.RegisterModel(new(models.User))
	orm.RunSyncdb("default", false, true)
}
