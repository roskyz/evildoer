package model

import "evildoer/protogen"

const prefix = "evildoer_v1_"

var StateMap = map[protogen.State]string{
	protogen.State_STATE_NORMAL:   "normal",
	protogen.State_STATE_DISABLED: "disabled",
	protogen.State_STATE_DELETED:  "deleted",
}

var StateMapReverse = map[string]protogen.State{
	"normal":   protogen.State_STATE_NORMAL,
	"disabled": protogen.State_STATE_DISABLED,
	"deleted":  protogen.State_STATE_DELETED,
}
