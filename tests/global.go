package tests

import (
	"path/filepath"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"runtime"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	// find app.conf path
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../../"+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)

	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlurls := beego.AppConfig.String("mysqlurls")
	mysqlport := beego.AppConfig.String("mysqlport")
	mysqldb := beego.AppConfig.String("mysqldb")

	beego.Warning(mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8mb4")
	//初始化数据库
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8mb4")
	orm.Debug = false
	//设置最大链接数
	orm.SetMaxIdleConns("default", 100)
	orm.SetMaxOpenConns("default", 300)
}
