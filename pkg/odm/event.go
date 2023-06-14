package odm

import (
	"context"
	"evildoer/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type eventInterface interface {
	GetEventByUUID(ctx context.Context, uuid string) (*model.Event, error)
	SearchEvents(ctx context.Context, filter bson.D, findOptions *options.FindOptions) (count int64, events []*model.Event, err error)
}

func (m *mongoBackend) GetEventByUUID(ctx context.Context, uuid string) (*model.Event, error) {
	var event model.Event
	if err := mgm.Coll(new(model.Event)).FindOne(ctx, bson.D{GenerateSingleFind("uuid", uuid)}).Decode(&event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (m *mongoBackend) SearchEvents(ctx context.Context, filter bson.D, findOptions *options.FindOptions) (count int64, events []*model.Event, err error) {
	events = make([]*model.Event, 0)
	if err := mgm.Coll(new(model.Event)).SimpleFindWithCtx(ctx, &events, filter, findOptions); err != nil {
		return 0, nil, err
	}
	count, err = mgm.Coll(new(model.Event)).CountDocuments(ctx, filter)
	return count, events, err
}
