package components

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	if beego.AppConfig.String("mysqlurls") != "" {
		mysqluser := beego.AppConfig.String("mysqluser")
		mysqlpass := beego.AppConfig.String("mysqlpass")
		mysqlurls := beego.AppConfig.String("mysqlurls")
		mysqlport := beego.AppConfig.String("mysqlport")
		mysqldb := beego.AppConfig.String("mysqldb")
		//初始化数据库
		orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlurls+":"+mysqlport+")/"+mysqldb+"?charset=utf8mb4")
		orm.Debug = true
		//设置最大链接数
		orm.SetMaxIdleConns("default", 100)
		orm.SetMaxOpenConns("default", 300)
	}
}