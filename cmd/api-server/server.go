package main

import (
	"github.com/astaxie/beego"
	"github.com/bullteam/zeus/pkg/components"
	_ "github.com/bullteam/zeus/pkg/routers"
	"github.com/spf13/cobra"
	"os"
)
var Args args
type args struct {
	ConfigFile     string
}
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
		Example: "zeus start -c conf/app.conf",
		RunE:    start,
	}
)

func init() {
	startCmd.PersistentFlags().StringVarP(&Args.ConfigFile, "config", "c", "conf/app.conf", "Start server with provided configuration file")
	rootCmd.AddCommand(startCmd)
}

func start(_ *cobra.Command, _ []string) error {
	beego.LoadAppConfig("ini", Args.ConfigFile)
	components.Init()//启动
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