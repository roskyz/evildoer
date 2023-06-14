package service

import (
	"context"
	"evildoer/protogen"

	"google.golang.org/grpc"
)

type subscriptionService struct {
	*protogen.UnimplementedSubscriptionSrvServer
}

func (ss *subscriptionService) ListSubscription(ctx context.Context, page *protogen.Page) (*protogen.ListSubscriptionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ss *subscriptionService) GetSubscription(ctx context.Context, response *protogen.IDReqOrResponse) (*protogen.SubscriptionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (ss *subscriptionService) CreateSubscription(ctx context.Context, subscription *protogen.Subscription) (*protogen.SubscriptionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func regSubscriptionService(grpcSrv *grpc.Server) error {
	protogen.RegisterSubscriptionSrvServer(grpcSrv, new(subscriptionService))
	return nil
}
