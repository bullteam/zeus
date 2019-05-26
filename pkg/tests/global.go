package tests

import (
	"github.com/astaxie/beego"
	"zeus/pkg/config"
	"zeus/pkg/dao"
)

var apppath string

func init() {
	// change to your application path
	apppath = "/data1/src/web/zeus"
	beego.TestBeegoInit(apppath)

	database, err := config.Database()
	if err != nil {
		beego.Error("failed to get database configuration: %v", err)
	}
	if err := dao.InitDatabase(database); err != nil {
		beego.Error("failed to initialize database: %v", err)
	}
}
