package dao

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/bullteam/zeus/pkg/models"

	"strconv"
	"strings"
	"sync"
)

// Database is an interface of different databases
type Database interface {
	// Name returns the name of database
	Name() string
	// String returns the details of database
	String() string
	// Register registers the database which will be used
	Register(alias ...string) error
}

// InitDatabase initializes the database
func InitDatabase(database *models.Database) error {
	db, err := getDatabase(database)
	if err != nil {
		return err
	}
	beego.Info("initializing database: %s", db.String())
	if err := db.Register(); err != nil {
		return err
	}
	beego.Info("initialize database completed")
	return nil
}

func getDatabase(database *models.Database) (db Database, err error) {
	db = NewMySQL(database.MySQL.Host,
		strconv.Itoa(database.MySQL.Port),
		database.MySQL.Username,
		database.MySQL.Password,
		database.MySQL.Database)
	return
}

var globalOrm orm.Ormer
var once sync.Once

// GetOrmer :set ormer singleton
func GetOrmer() orm.Ormer {
	once.Do(func() {
		globalOrm = orm.NewOrm()
	})
	return globalOrm
}

func paginateForRawSQL(sql string, limit, offset int64) string {
	return fmt.Sprintf("%s limit %d offset %d", sql, limit, offset)
}

func escape(str string) string {
	str = strings.Replace(str, `%`, `\%`, -1)
	str = strings.Replace(str, `_`, `\_`, -1)
	return str
}
