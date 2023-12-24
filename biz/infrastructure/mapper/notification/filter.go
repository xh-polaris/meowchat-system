package notification

import (
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/consts"
)

type FilterOptions struct {
	OnlyUserId     *string
	OnlyType       *system.NotificationType
	OnlyTargetType *system.NotificationTargetType
	FirstId        *primitive.ObjectID
	LastId         *primitive.ObjectID
	IsRead         *bool
}

type MongoFilter struct {
	m bson.M
	*FilterOptions
}

func makeMongoFilter(options *FilterOptions) bson.M {
	return (&MongoFilter{
		m:             bson.M{},
		FilterOptions: options,
	}).toBson()
}

func (f *MongoFilter) toBson() bson.M {
	f.CheckOnlyUserId()
	f.CheckOnlyType()
	f.CheckOnlyTargetType()
	f.CheckRange()
	f.CheckOnlyIsRead()
	return f.m
}

func (f *MongoFilter) CheckOnlyUserId() {
	if f.OnlyUserId != nil {
		f.m[consts.TargetUserId] = *f.OnlyUserId
	}
}

func (f *MongoFilter) CheckOnlyType() {
	if f.OnlyType != nil {
		f.m[consts.Type] = *f.OnlyType
	}
}

func (f *MongoFilter) CheckOnlyTargetType() {
	if f.OnlyTargetType != nil {
		f.m[consts.TargetType] = *f.OnlyTargetType
	}
}

func (f *MongoFilter) CheckRange() {
	if f.LastId != nil && f.FirstId != nil {
		f.m[consts.NotificationId] = bson.M{"$gte": f.LastId, "$lte": f.FirstId}
	}
}

func (f *MongoFilter) CheckOnlyIsRead() {
	if f.IsRead != nil {
		f.m[consts.IsRead] = bson.M{"$exists": *f.IsRead}
	}
}
