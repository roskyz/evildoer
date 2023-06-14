package model

import (
	"evildoer/protogen"

	"github.com/kamva/mgm/v3"
)

type Event struct {
	mgm.DefaultModel `bson:",inline"`
	AccessMode       string                 `json:"access_mode" bson:"access_mode"`
	Form             string                 `json:"form_key" bson:"form_key"`
	FormVersion      int32                  `json:"form_version" bson:"form_version"`
	UUID             string                 `json:"uuid" bson:"uuid"`
	Content          map[string]interface{} `json:"content" bson:"content"`
	CloudEvents      map[string]interface{} `json:"cloudevents" bson:"cloudevents"`
	Creator          string                 `json:"creator" bson:"creator"`
}

func (Event) CollectionName() string {
	return prefix + "event"
}

func (e Event) Conv() *protogen.Event {
	return &protogen.Event{
		Uuid:        e.UUID,
		AccessMode:  e.AccessMode,
		FormKey:     e.Form,
		FormVersion: e.FormVersion,
		InputData:   marshaler(e.Content),
		Cloudevents: marshaler(e.CloudEvents),
		Creator:     e.Creator,
	}
}

func (Form) NewEvent(event *protogen.Event) *Event {
	return &Event{
		DefaultModel: mgm.DefaultModel{},
		AccessMode:   event.AccessMode,
		Form:         event.FormKey,
		FormVersion:  event.FormVersion,
		UUID:         event.Uuid,
		Content:      unmarshaler(event.InputData),
		CloudEvents:  unmarshaler(event.Cloudevents),
		Creator:      event.Creator,
	}
}
