package client

import (
	"context"
	"evildoer/pkg/authmgr"
	"evildoer/protogen"

	"google.golang.org/grpc"
)

func NewGroupClient(ctx context.Context, targetURI string, auth *authmgr.AuthOpt) (protogen.GroupSrvClient, error) {
	var conn grpc.ClientConnInterface
	var err error
	if conn, err = newCaller(ctx, targetURI, WithAuthInterceptor(auth)); err != nil {
		return nil, err
	}
	return newGroupClient(conn), nil
}

func newGroupClient(cc grpc.ClientConnInterface) protogen.GroupSrvClient {
	return protogen.NewGroupSrvClient(cc)
}
