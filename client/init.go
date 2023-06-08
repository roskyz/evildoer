package client

import (
	"context"
	"evildoer/pkg/authmgr"
	"evildoer/protogen"
)

type Client interface {
	protogen.GroupSrvClient
}

type clientImpl struct {
	protogen.GroupSrvClient
}

func NewAggrClient(ctx context.Context, targetURI string, auth *authmgr.AuthOpt) (Client, error) {
	conn, err := newCaller(ctx, targetURI, WithAuthInterceptor(auth))
	if err != nil {
		return nil, err
	}
	client := &clientImpl{
		GroupSrvClient: newGroupClient(conn),
	}
	return client, nil
}
