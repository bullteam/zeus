package main

import (
	_ "github.com/bullteam/zeus/pkg/routers"
	"github.com/astaxie/beego"
)
func main() {
	//st := flag.String("s","http","Type of serve,including http,grpc")
	//pt := flag.String("p","8188","port of grpc server")
	//flag.Parse()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
