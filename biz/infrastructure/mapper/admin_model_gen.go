// Code generated by goctl. DO NOT EDIT.
package mapper

import (
	"context"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var prefixAdminCacheKey = "cache:admin:"

type adminModel interface {
	Insert(ctx context.Context, data *db.Admin) error
	FindOne(ctx context.Context, id string) (*db.Admin, error)
	Update(ctx context.Context, data *db.Admin) error
	Delete(ctx context.Context, id string) error
}

type defaultAdminModel struct {
	conn *monc.Model
}

func newDefaultAdminModel(conn *monc.Model) *defaultAdminModel {
	return &defaultAdminModel{conn: conn}
}

func (m *defaultAdminModel) Insert(ctx context.Context, data *db.Admin) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixAdminCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *defaultAdminModel) FindOne(ctx context.Context, id string) (*db.Admin, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data db.Admin
	key := prefixAdminCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminModel) Update(ctx context.Context, data *db.Admin) error {
	data.UpdateAt = time.Now()
	key := prefixAdminCacheKey + data.ID.Hex()
	_, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return err
}

func (m *defaultAdminModel) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}
	key := prefixAdminCacheKey + id
	_, err = m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return err
}
