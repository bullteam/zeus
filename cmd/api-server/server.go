package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/bullteam/zeus/pkg/routers"
)

func main() {
	usage()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func usage() {
	fmt.Printf("%s\n", usageStr)
}

var usageStr = `
  ______              
 |___  /              
    / / ___ _   _ ___ 
   / / / _ \ | | / __|
  / /_|  __/ |_| \__ \
 /_____\___|\__,_|___/
`
