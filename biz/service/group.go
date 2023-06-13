package service

import (
	context "context"
	"evildoer/model"
	"evildoer/pkg/odm"
	"evildoer/protogen"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type groupService struct {
	*protogen.UnimplementedGroupSrvServer
}

func (gs *groupService) ListGroup(ctx context.Context, req *protogen.ListGroupReq) (*protogen.ListGroupResponse, error) {
	page := req.GetPage()
	searchOption := req.GetSearchOption()

	var filter bson.D
	if content := searchOption.GetContent(); content != "" {
		filter = odm.GenerateSearch([]*odm.Field{
			{Name: "name", Content: content},
			{Name: "description", Content: content},
		})
	}

	count, groups, err := odm.GetMongo().ListGroups(ctx, filter,
		odm.GenerateFindOption(searchOption.GetSort().GetField(), searchOption.GetSort().GetAsc(), page.GetOffset(), page.GetLimit()),
	)

	response := &protogen.ListGroupResponse{
		DefaultResponse: defaultResponse(err),
		Count:           count,
	}
	for _, group := range groups {
		response.Groups = append(response.Groups, group.Conv())
	}
	return response, nil
}

func (gs *groupService) GetGroup(ctx context.Context, response *protogen.IDReqOrResponse) (*protogen.GetGroupResponse, error) {
	group, err := odm.GetMongo().GetGroup(ctx, string(response.GetId()))
	return &protogen.GetGroupResponse{
		DefaultResponse: defaultResponse(err),
		Group:           group.Conv(),
	}, nil
}

func (gs *groupService) CreateGroup(ctx context.Context, group *protogen.Group) (*protogen.CreateGroupResponse, error) {
	createGroup := new(model.Group).NewGroup(group)
	err := odm.GetMongo().CreateGroup(ctx, createGroup)
	return &protogen.CreateGroupResponse{
		DefaultResponse: defaultResponse(err),
		IdResp:          &protogen.IDReqOrResponse{Id: []byte(createGroup.ID.Hex())},
	}, nil
}

func (gs *groupService) UpdateGroup(ctx context.Context, group *protogen.Group) (*protogen.DefaultResponse, error) {
	updateGroup, err := odm.GetMongo().GetGroup(ctx, string(group.GetId()))
	if err == nil {
		err = odm.GetMongo().UpdateGroup(ctx, updateGroup.UpdateGroup(group))
	}
	return defaultResponse(err), nil
}

func (gs *groupService) ListMode(ctx context.Context, empty *emptypb.Empty) (*protogen.ListModeResponse, error) {
	listModeResponse := &protogen.ListModeResponse{
		DefaultResponse: defaultResponse(),
		GroupState: []protogen.State{
			protogen.State_STATE_NORMAL,
			protogen.State_STATE_DISABLED,
			protogen.State_STATE_DELETED,
		},
	}
	return listModeResponse, nil
}

func regGroupService(grpcSrv *grpc.Server) error {
	protogen.RegisterGroupSrvServer(grpcSrv, new(groupService))
	return nil
}
