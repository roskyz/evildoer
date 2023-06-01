package service

import (
	"google.golang.org/grpc"
)

func NewService() (*grpc.Server, error) {
	grpcSrv := grpc.NewServer()
	// todo: errgroup
	if err := regGroupService(grpcSrv); err != nil {
		return nil, err
	}
	return grpcSrv, nil
}
