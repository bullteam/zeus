package main

import (
	"context"
	pb "zeus/cmd/grpc-linker/grpc-server/proto"
	"zeus/pkg/service"
)

type ApiAuthServer struct{}

func (aas *ApiAuthServer) CheckPerm(ctx context.Context,request *pb.PermRequest) (*pb.PermResponse,error){
	permService := service.PermService{}
	return &pb.PermResponse{Pass:permService.CheckPermByUid(int(request.Uid),request.Perm,request.Domain)},nil
}
