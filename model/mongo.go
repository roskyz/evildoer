package model

import (
	"encoding/json"
	"evildoer/protogen"
)

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

var marshaler = func(data interface{}) []byte {
	buffer, _ := json.Marshal(data)
	return buffer
}

var unmarshaler = func(buffer []byte) map[string]interface{} {
	var result = make(map[string]interface{})
	_ = json.Unmarshal(buffer, &result)
	return result
}
