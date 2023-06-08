package odm

import (
	"context"
	"evildoer/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type groupInterface interface {
	ListGroups(ctx context.Context, filter bson.D, findOptions *options.FindOptions) (count int64, groups []*model.Group, err error)
	GetGroup(ctx context.Context, id string) (group *model.Group, err error)
	CreateGroup(ctx context.Context, group *model.Group) error
	UpdateGroup(ctx context.Context, group *model.Group) error
}

func (m *mongoBackend) ListGroups(ctx context.Context, filter bson.D, findOptions *options.FindOptions) (count int64, groups []*model.Group, err error) {
	groups = make([]*model.Group, 0)
	err = mgm.Coll(new(model.Group)).SimpleFindWithCtx(ctx, &groups, filter, findOptions)
	if err != nil {
		return 0, nil, err
	}
	count, err = mgm.Coll(new(model.Group)).CountDocuments(ctx, filter)
	return count, groups, err
}

func (m *mongoBackend) GetGroup(ctx context.Context, id string) (*model.Group, error) {
	var group model.Group
	if err := mgm.Coll(new(model.Group)).FindByIDWithCtx(ctx, id, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

func (m *mongoBackend) CreateGroup(ctx context.Context, group *model.Group) error {
	return mgm.Coll(new(model.Group)).CreateWithCtx(ctx, group)
}

func (m *mongoBackend) UpdateGroup(ctx context.Context, group *model.Group) error {
	return mgm.Coll(new(model.Group)).UpdateWithCtx(ctx, group)
}
