package service

import (
	"context"
	"evildoer/model"
	"evildoer/pkg/odm"
	"evildoer/protogen"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type formService struct {
	*protogen.UnimplementedFormSrvServer
}

func (fs *formService) ListForm(ctx context.Context, req *protogen.ListFormReq) (*protogen.ListFormResponse, error) {
	page := req.GetPage()
	searchOption := req.GetFormSearch().GetSearchOption()

	var filter bson.D
	if content := searchOption.GetContent(); content != "" {
		filter = odm.GenerateSearch([]*odm.Field{
			{Name: "name", Content: content},
			{Name: "description", Content: content},
		})
	}

	sortField := searchOption.GetSort().GetField()
	sortAsc := searchOption.GetSort().GetAsc()

	latestVersion := req.GetFormSearch().GetLatest()
	if latestVersion && sortField != "created_at" && !sortAsc {
		filter = append(filter, odm.GenerateSingleFind("latest", true))
	}

	count, forms, err := odm.GetMongo().ListForms(ctx, filter,
		odm.GenerateFindOption(sortField, sortAsc, page.GetOffset(), page.GetLimit()),
	)

	response := &protogen.ListFormResponse{
		DefaultResponse: defaultResponse(err),
		Count:           count,
	}
	for _, form := range forms {
		response.Forms = append(response.Forms, form.Conv())
	}
	return response, nil
}

func (fs *formService) GetForm(ctx context.Context, response *protogen.IDReqOrResponse) (*protogen.GetFormResponse, error) {
	form, err := odm.GetMongo().GetForm(ctx, string(response.GetId()))
	return &protogen.GetFormResponse{
		DefaultResponse: defaultResponse(err),
		Form:            form.Conv(),
	}, nil
}

func (fs *formService) ChangeState(ctx context.Context, req *protogen.ChangeStateReq) (*protogen.DefaultResponse, error) {
	return defaultResponse(odm.GetMongo().UpdateFormState(
		ctx, string(req.GetIdReq().GetId()), model.StateMap[req.GetState()])), nil
}

func (fs *formService) CreateForm(ctx context.Context, form *protogen.Form) (*protogen.FormAndGroupResponse, error) {
	createForm := new(model.Form).NewForm(form)
	err := odm.GetMongo().CreateForm(ctx, createForm)
	group, form, err := fs.genFormResponse(ctx, createForm)
	// todo return nil would be manutal? caller invoker handler response err not response
	return &protogen.FormAndGroupResponse{
		DefaultResponse: defaultResponse(err),
		Group:           group,
		Form:            form,
	}, nil
}

func (fs *formService) genFormResponse(ctx context.Context, form *model.Form) (*protogen.Group, *protogen.Form, error) {
	group, err := odm.GetMongo().GetGroupByKey(ctx, form.Key)
	if err != nil {
		return nil, nil, err
	}
	return group.Conv(), form.Conv(), nil
}

func (fs *formService) UpdateForm(ctx context.Context, req *protogen.UpdateFormReq) (*protogen.FormAndGroupResponse, error) {
	oldForm, err := odm.GetMongo().GetForm(ctx, string(req.GetIdReq().GetId()))
	if err != nil {
		return &protogen.FormAndGroupResponse{
			DefaultResponse: defaultResponse(err),
			Group:           nil,
			Form:            nil,
		}, nil
	}
	oldForm.Latest = false
	newForm := new(model.Form).NewForm(req.GetForm())
	err = odm.GetMongo().Transaction(ctx, func(sessionCtx context.Context) error {
		if err := odm.GetMongo().UpdateForm(sessionCtx, oldForm); err != nil {
			return err
		}
		if err := odm.GetMongo().CreateForm(sessionCtx, newForm); err != nil {
			return err
		}
		return nil
	})
	group, form, err := fs.genFormResponse(ctx, newForm)
	return &protogen.FormAndGroupResponse{
		DefaultResponse: defaultResponse(err),
		Group:           group,
		Form:            form,
	}, nil
}

func regFormService(grpcSrv *grpc.Server) error {
	protogen.RegisterFormSrvServer(grpcSrv, new(formService))
	return nil
}
