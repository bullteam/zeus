package dao

import (
	"fmt"
	"github.com/bullteam/zeus/pkg/utils"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //register mysql driver
)

type mysql struct {
	host     string
	port     string
	usr      string
	pwd      string
	database string
}

// NewMySQL returns an instance of mysql
func NewMySQL(host, port, usr, pwd, database string) Database {
	return &mysql{
		host:     host,
		port:     port,
		usr:      usr,
		pwd:      pwd,
		database: database,
	}
}

// Register registers MySQL as the underlying database used
func (m *mysql) Register(alias ...string) error {
	if err := utils.TestTCPConn(m.host+":"+m.port, 60, 2); err != nil {
		return err
	}
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		return err
	}

	an := "default"
	if len(alias) != 0 {
		an = alias[0]
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=%s", m.usr, m.pwd, m.host, m.port, m.database, "Asia%2FHarbin")
	return orm.RegisterDataBase(an, "mysql", conn)
}

// Name returns the name of MySQL
func (m *mysql) Name() string {
	return "MySQL"
}

// String returns the details of database
func (m *mysql) String() string {
	return fmt.Sprintf("type-%s host-%s port-%s user-%s database-%s",
		m.Name(), m.host, m.port, m.usr, m.database)
}
