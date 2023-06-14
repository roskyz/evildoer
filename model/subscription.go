package model

import "github.com/kamva/mgm/v3"

type Subscription struct {
	mgm.DefaultModel `bson:",inline"`
	Key              string             `json:"key" bson:"key"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	RuleGroup        map[string][]*Rule `json:"rule_group" bson:"rule_group"`
	ConsumerGroup    []*Consumer        `json:"consumer_group" bson:"consumer_group"`
	State            string             `json:"state" bson:"state"`
	Creator          string             `json:"creator" bson:"creator"`
	FormKey          string             `json:"form_key" bson:"form_key"`
}

type Rule struct {
	Pattern string      `json:"pattern" bson:"pattern"`
	Field   string      `json:"field" bson:"field"`
	Value   interface{} `json:"value" bson:"value"`
}

type Consumer struct {
	Type   string                 `json:"type" bson:"type"`
	Fields map[string]interface{} `json:"fields" bson:"fields"`
}

func (Subscription) CollectionName() string {
	return prefix + "subscription"
}
