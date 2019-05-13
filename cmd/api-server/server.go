package api_server

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"zeus/pkg/components"
	"zeus/pkg/config"
	"zeus/pkg/dao"
	_ "zeus/pkg/routers"
	"github.com/spf13/cobra"
)

var (
	StartCmd = &cobra.Command{
		Use:     "start",
		Short:   "Start zeus API server",
		Example: "zeus start -c ./conf",
		RunE:    start,
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&components.Args.ConfigFile, "config", "c", "./conf", "Start server with provided configuration file")
}

func start(_ *cobra.Command, _ []string) error {
	_ = beego.LoadAppConfig("ini", components.Args.ConfigFile+"/app.conf")
	_ = beego.AddFuncMap("i18n", i18n.Tr)
	components.RedisInit()
	usae()
	database, err := config.Database()
	if err != nil {
		beego.Error("failed to get database configuration: %v", err)
	}
	if err := dao.InitDatabase(database); err != nil {
		beego.Error("failed to initialize database: %v", err)
	}
	beego.Run()
	return nil
}
func usae() {
	usageStr := `
  ______              
 |___  /              
    / / ___ _   _ ___ 
   / / / _ \ | | / __|
  / /_|  __/ |_| \__ \
 /_____\___|\__,_|___/
`
	fmt.Printf("%s\n", usageStr)
}
