// Code generated by goctl. DO NOT EDIT.
package user_role

import (
	"context"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper"

	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var prefixUserRoleCacheKey = "cache:userRole:"

type userRoleModel interface {
	Insert(ctx context.Context, data *db.UserRole) error
	FindOne(ctx context.Context, id string) (*db.UserRole, error)
	Update(ctx context.Context, data *db.UserRole) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultUserRoleModel struct {
	conn *monc.Model
}

func newDefaultUserRoleModel(conn *monc.Model) *defaultUserRoleModel {
	return &defaultUserRoleModel{conn: conn}
}

func (m *defaultUserRoleModel) Insert(ctx context.Context, data *db.UserRole) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixUserRoleCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultUserRoleModel) FindOne(ctx context.Context, id string) (*db.UserRole, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, mapper.ErrInvalidObjectId
	}

	var data db.UserRole
	key := prefixUserRoleCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, mapper.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserRoleModel) Update(ctx context.Context, data *db.UserRole) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixUserRoleCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultUserRoleModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, mapper.ErrInvalidObjectId
	}
	key := prefixUserRoleCacheKey + id
	res, err := m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return res, err
}
