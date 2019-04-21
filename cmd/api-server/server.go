package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bullteam/zeus/pkg/components"
	"github.com/bullteam/zeus/pkg/config"
	"github.com/bullteam/zeus/pkg/dao"
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
	components.RedisInit()
	usae()
	database, err := config.Database()
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
func usae(){
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

