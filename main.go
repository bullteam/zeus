package main

import (
	_ "github.com/bullteam/zeus/routers"
	"github.com/astaxie/beego"
	//_ "github.com/astaxie/beego/cache/redis"
	pb "github.com/bullteam/zeus/grpc-server/proto"
	"flag"
	"google.golang.org/grpc/reflection"
	"net"
	"log"
	"google.golang.org/grpc"
)
func init(){

}
func main() {
	st := flag.String("s","http","Type of serve,including http,grpc")
	pt := flag.String("p","8188","port of grpc server")
	flag.Parse()
	if *st != "grpc"{
		if beego.BConfig.RunMode == "dev" {
			beego.BConfig.WebConfig.DirectoryIndex = true
			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		}
		beego.Run()
	}else{
		lis, err := net.Listen("tcp", ":"+*pt)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}else{
			log.Printf("Grpc serve on port:"+*pt)
		}
		s := grpc.NewServer()
		pb.RegisterApiauthServer(s,&ApiAuthServer{})
		//pb.RegisterCasbinServer(s, server.NewServer())
		// Register reflection service on gRPC server.
		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
