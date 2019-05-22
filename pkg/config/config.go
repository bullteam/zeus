package config

import (
	"github.com/astaxie/beego"
	"zeus/pkg/models"
)

// Database returns database settings
func Database() (*models.Database, error) {
	database := &models.Database{}
	mysql := &models.MySQL{}
	mysql.Host = beego.AppConfig.String("mysqlurls")
	port, err := beego.AppConfig.Int("mysqlport")
	if err != nil {
		return nil, err
	}
	mysql.Port = port
	mysql.Username = beego.AppConfig.String("mysqluser")
	mysql.Password = beego.AppConfig.String("mysqlpass")
	mysql.Database = beego.AppConfig.String("mysqldb")
	database.MySQL = mysql
	return database, nil
}
