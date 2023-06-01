package service

import (
	"evildoer/protogen"

	"google.golang.org/grpc"
)

type groupService struct {
	*protogen.UnimplementedGroupSrvServer
}

func regGroupService(grpcSrv *grpc.Server) error {
	protogen.RegisterGroupSrvServer(grpcSrv, new(groupService))
	return nil
}
