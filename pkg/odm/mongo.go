package odm

import (
	"evildoer/conf"
	"evildoer/utils"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BizInterfaces interface {
	groupInterface
}

var backendGetter utils.OnceGetter[BizInterfaces] = utils.MakeGetter(func() BizInterfaces {
	mongoCfg := conf.GetConf().Mongo
	utils.AssertErr(mgm.SetDefaultConfig(
		&mgm.Config{CtxTimeout: time.Second * 3}, mongoCfg.Database,
		options.Client().
			ApplyURI(mongoCfg.Address).
			SetAuth(options.Credential{
				Username: mongoCfg.Username,
				Password: mongoCfg.Password,
			})))
	return new(mongoBackend)
})

type mongoBackend struct{}

func GetMongo() BizInterfaces {
	return backendGetter.Getter()
}

type Field struct {
	Name, Content string
}

func GenerateSearch(fields []*Field) bson.D {
	var filter bson.D
	for _, field := range fields {
		filter = append(filter, generateSingleRegex(field.Name, field.Content))
	}
	return filter
}

func generateSingleRegex(fieldName string, value string) bson.E {
	return bson.E{Key: fieldName, Value: primitive.Regex{Pattern: value}}
}

func GenerateSingleFind(fieldName string, value interface{}) bson.D {
	return bson.D{bson.E{Key: fieldName, Value: value}}
}

func GenerateFindOption(sortField string, asc bool, offset, limit int64) *options.FindOptions {
	option := new(options.FindOptions)
	option = option.SetSkip(offset).SetLimit(limit)

	if sortField != "" {
		var sortOpt bson.D
		var order = 1
		if !asc {
			order = -1
		}
		sortOpt = append(sortOpt, bson.E{Key: sortField, Value: order})
		option = option.SetSort(sortField)
	}

	return option
}
