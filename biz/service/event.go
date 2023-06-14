package service

import (
	"context"
	"evildoer/pkg/odm"
	"evildoer/protogen"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type eventService struct {
	*protogen.UnimplementedEventSrvServer
}

func (es *eventService) SearchEvent(ctx context.Context, req *protogen.SearchEventReq) (*protogen.SearchEventResponse, error) {
	// todo changed to elastic
	page := req.GetPage()
	eventSearchOpt := req.GetEventSearch()
	// todo check all optional fields npe
	searchOption := eventSearchOpt.GetSearchOption()
	var filter bson.D
	if searchOption.GetCreatedAfter() != nil && searchOption.GetCreatedBefore() != nil {
		filter = append(filter,
			bson.E{Key: "created_at", Value: bson.M{"$gte": searchOption.GetCreatedAfter()}},
			bson.E{Key: "created_at", Value: bson.M{"$lte": searchOption.GetCreatedBefore()}},
		)
	}

	// todo created_at, creator ... fields can all use parser
	if eventSearchOpt.GetFormKey() != "" {
		filter = append(filter, odm.GenerateSingleFind("form_key", eventSearchOpt.GetFormKey()))
	}

	if eventSearchOpt.GetCreator() != "" {
		filter = append(filter, odm.GenerateSingleFind("creator", eventSearchOpt.GetCreator()))
	}

	count, events, err := odm.GetMongo().SearchEvents(ctx, filter,
		odm.GenerateFindOption(searchOption.GetSort().GetField(), searchOption.GetSort().GetAsc(), page.GetOffset(), page.GetLimit()),
	)
	response := &protogen.SearchEventResponse{
		DefaultResponse: defaultResponse(err),
		Count:           count,
	}
	for _, event := range events {
		response.Events = append(response.Events, event.Conv())
	}
	return response, nil
}

func (es *eventService) GetEvent(ctx context.Context, req *protogen.GetEventReq) (*protogen.GetEventResponse, error) {
	event, err := odm.GetMongo().GetEventByUUID(ctx, req.GetUuid())
	return &protogen.GetEventResponse{
		DefaultResponse: defaultResponse(err),
		Event:           event.Conv(),
	}, nil
}

func (es *eventService) CreateEvent(ctx context.Context, req *protogen.CreateEventReq) (*protogen.Event, error) {
	//eventID := uuid.NewString()
	return nil, nil
}

func regEventService(grpcSrv *grpc.Server) error {
	protogen.RegisterEventSrvServer(grpcSrv, new(eventService))
	return nil
}
