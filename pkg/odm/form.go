package odm

import (
	"context"
	"evildoer/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type formInterface interface {
	ListForms(ctx context.Context, filter bson.D, findOptions *options.FindOptions) (count int64, forms []*model.Form, err error)
	GetForm(ctx context.Context, id string) (form *model.Form, err error)
	UpdateForm(ctx context.Context, form *model.Form) error
	UpdateFormState(ctx context.Context, id string, state string) error
	CreateForm(ctx context.Context, form *model.Form) error
}

// ListForms sort by created_at so all forms would be the latest version
func (m *mongoBackend) ListForms(ctx context.Context, filter bson.D, findOptions *options.FindOptions) (count int64, forms []*model.Form, err error) {
	forms = make([]*model.Form, 0)
	err = mgm.Coll(new(model.Form)).SimpleFindWithCtx(ctx, &forms, filter, findOptions)
	if err != nil {
		return 0, nil, err
	}
	count, err = mgm.Coll(new(model.Form)).CountDocuments(ctx, filter)
	return count, forms, err
}

func (m *mongoBackend) GetForm(ctx context.Context, id string) (*model.Form, error) {
	var form model.Form
	if err := mgm.Coll(new(model.Form)).FindByIDWithCtx(ctx, id, &form); err != nil {
		return nil, err
	}
	return &form, nil
}

func (m *mongoBackend) UpdateForm(ctx context.Context, form *model.Form) error {
	return mgm.Coll(new(model.Form)).UpdateWithCtx(ctx, form)
}

func (m *mongoBackend) UpdateFormState(ctx context.Context, id string, state string) error {
	update := bson.D{{"$set", bson.D{{"state", state}}}}
	_, err := mgm.Coll(new(model.Form)).UpdateByID(ctx, id, update, options.Update().SetUpsert(false))
	return err
}

func (m *mongoBackend) CreateForm(ctx context.Context, form *model.Form) error {
	return mgm.Coll(new(model.Form)).CreateWithCtx(ctx, form)
}
