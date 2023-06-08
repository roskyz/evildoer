package model

import (
	"evildoer/protogen"

	"github.com/kamva/mgm/v3"
)

type Group struct {
	mgm.DefaultModel `bson:",inline"`
	Key              string   `json:"key" bson:"key"`
	Name             string   `json:"name" bson:"name"`
	Description      string   `json:"description" bson:"description"`
	State            string   `json:"state" bson:"state"`
	Creator          string   `json:"creator" bson:"creator"`
	AccessMode       []string `json:"access_mode" bson:"access_mode"`
}

func (Group) CollectionName() string {
	return prefix + "group"
}

var groupStateMap = map[protogen.GroupState]string{
	protogen.GroupState_GROUP_STATE_NORMAL:   "normal",
	protogen.GroupState_GROUP_STATE_DISABLED: "disabled",
	protogen.GroupState_GROUP_STATE_DELETED:  "deleted",
}

var groupStateMapReverse = map[string]protogen.GroupState{
	"normal":   protogen.GroupState_GROUP_STATE_NORMAL,
	"disabled": protogen.GroupState_GROUP_STATE_DISABLED,
	"deleted":  protogen.GroupState_GROUP_STATE_DELETED,
}

var accessModeMap = map[protogen.AccessMode]string{
	protogen.AccessMode_ACCESS_MODE_API: "api",
	protogen.AccessMode_ACCESS_MODE_APP: "app",
	protogen.AccessMode_ACCESS_MODE_MQ:  "mq",
}

var accessModeMapReverse = map[string]protogen.AccessMode{
	"api": protogen.AccessMode_ACCESS_MODE_API,
	"app": protogen.AccessMode_ACCESS_MODE_APP,
	"mq":  protogen.AccessMode_ACCESS_MODE_MQ,
}

func (g Group) Conv() *protogen.Group {
	var accessModes = make([]protogen.AccessMode, 0, len(g.AccessMode))
	for _, mode := range g.AccessMode {
		accessModes = append(accessModes, accessModeMapReverse[mode])
	}
	return &protogen.Group{
		Id:          []byte(g.ID.Hex()),
		Key:         g.Key,
		Name:        g.Name,
		Description: g.Description,
		State:       groupStateMapReverse[g.State],
		Creator:     g.Creator,
		AccessMode:  accessModes,
	}
}

func (Group) NewGroup(group *protogen.Group) *Group {
	var accessModes = make([]string, 0, len(group.AccessMode))
	for _, mode := range group.AccessMode {
		accessModes = append(accessModes, accessModeMap[mode])
	}
	return &Group{
		Key:         group.Key,
		Name:        group.Name,
		Description: group.Description,
		State:       groupStateMap[group.State],
		Creator:     group.Creator,
		AccessMode:  accessModes,
	}
}

func (g Group) UpdateGroup(group *protogen.Group) *Group {
	updateGroup := g.NewGroup(group)
	updateGroup.DefaultModel = g.DefaultModel
	return updateGroup
}
