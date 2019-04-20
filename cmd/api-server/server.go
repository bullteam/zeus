package main

import (
	"github.com/astaxie/beego"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/dao"
	"github.com/bullteam/zeus/pkg/models"
	_ "github.com/bullteam/zeus/pkg/routers"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:               "zeus",
		Short:             "zeus API server",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long:              `Start zeus API server`,
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	}
	startCmd = &cobra.Command{
		Use:     "start",
		Short:   "Start zeus API server",
		Example: "zeus start -c ./conf",
		RunE:    start,
	}
)

func init() {
	startCmd.PersistentFlags().StringVarP(&components.Args.ConfigFile, "config", "c", "./conf", "Start server with provided configuration file")
	rootCmd.AddCommand(startCmd)
}

func start(_ *cobra.Command, _ []string) error {
	beego.LoadAppConfig("ini", components.Args.ConfigFile+"/app.conf")
	components.Init()//启动
	database, err := Database()
	if err != nil {
		beego.Error("failed to get database configuration: %v", err)
	}

	if err := dao.InitDatabase(database); err != nil {
		beego.Error("failed to initialize database: %v", err)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

// Database returns database settings
func Database() (*models.Database, error) {
	database := &models.Database{}
	mysql := &models.MySQL{}
	mysql.Host = beego.AppConfig.String("mysqlurls")
	port,err :=beego.AppConfig.Int("mysqlport")
	if err != nil {
		return nil,err
	}
	mysql.Port = port
	mysql.Username = beego.AppConfig.String("mysqluser")
	mysql.Password = beego.AppConfig.String("mysqlpass")
	mysql.Database = beego.AppConfig.String("mysqldb")
	database.MySQL = mysql
	return database, nil
}