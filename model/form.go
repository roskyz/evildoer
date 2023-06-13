package model

import (
	"encoding/json"
	"evildoer/protogen"

	"github.com/kamva/mgm/v3"
)

type Form struct {
	mgm.DefaultModel `bson:",inline"`
	Key              string                 `json:"key" bson:"key"`
	Name             string                 `json:"name" bson:"name"`
	Description      string                 `json:"description" bson:"description"`
	State            string                 `json:"state" bson:"state"`
	Creator          string                 `json:"creator" bson:"creator"` // todo Group Form set creator
	Latest           bool                   `json:"latest" bson:"latest"`
	Version          int32                  `json:"version" bson:"version"`
	Group            string                 `json:"group" bson:"group"`
	ContentScheme    map[string]interface{} `json:"content_scheme" bson:"content_scheme"`
	ContentSample    map[string]interface{} `json:"content_sample" bson:"content_sample"`
}

func (Form) CollectionName() string {
	return prefix + "form"
}

func (f Form) Conv() *protogen.Form {
	marshaler := func(data interface{}) []byte {
		buffer, _ := json.Marshal(data)
		return buffer
	}
	return &protogen.Form{
		Id:            []byte(f.ID.Hex()),
		Key:           f.Key,
		Name:          f.Name,
		Description:   f.Description,
		GroupKey:      f.Group,
		Version:       f.Version,
		LatestVersion: f.Latest,
		ContentScheme: marshaler(f.ContentScheme),
		ContentSample: marshaler(f.ContentSample),
		State:         StateMapReverse[f.State],
		Creator:       f.Creator,
	}
}

func (Form) NewForm(form *protogen.Form) *Form {
	unmarsaler := func(buffer []byte) map[string]interface{} {
		var result = make(map[string]interface{})
		_ = json.Unmarshal(buffer, &result)
		return result
	}
	return &Form{
		DefaultModel:  mgm.DefaultModel{},
		Key:           form.Key,
		Name:          form.Name,
		Description:   form.Description,
		State:         StateMap[form.State],
		Creator:       form.Creator,
		Latest:        true,
		Version:       form.Version,
		Group:         form.GroupKey,
		ContentScheme: unmarsaler(form.ContentScheme),
		ContentSample: unmarsaler(form.ContentSample),
	}
}
