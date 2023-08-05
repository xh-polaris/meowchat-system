package mapper

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"meowchat-system/biz/infrastructure/config"
	"meowchat-system/biz/infrastructure/data/db"
)

const AdminCollectionName = "admin"

var _ AdminModel = (*customAdminModel)(nil)

type (
	// AdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminModel.
	AdminModel interface {
		adminModel
		ListAdmin(ctx context.Context, query *system.ListAdminReq) ([]*db.Admin, int64, error)
		UpdateAdmin(ctx context.Context, req *system.UpdateAdminReq) error
	}

	customAdminModel struct {
		*defaultAdminModel
	}
)

func (m customAdminModel) UpdateAdmin(ctx context.Context, req *system.UpdateAdminReq) error {
	key := prefixAdminCacheKey + req.Id

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return ErrInvalidObjectId
	}

	filter := bson.M{
		"_id": oid,
	}
	set := bson.M{
		"communityId": req.CommunityId,
		"name":        req.Name,
		"title":       req.Title,
		"phone":       req.Phone,
		"wechat":      req.Wechat,
		"avatarUrl":   req.AvatarUrl,
	}

	// 更新数据
	update := bson.M{
		"$set": set,
	}

	option := options.UpdateOptions{}
	option.SetUpsert(false)

	_, err = m.conn.UpdateOne(ctx, key, filter, update, &option)
	return err
}

func (m customAdminModel) ListAdmin(ctx context.Context, req *system.ListAdminReq) ([]*db.Admin, int64, error) {
	var resp []*db.Admin

	filter := bson.M{
		"communityId": req.CommunityId,
	}
	findOptions := ToFindOptions(req.Page, req.PageSize, req.Sort)

	err := m.conn.Find(ctx, &resp, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}

	count, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return resp, count, nil
}

// NewAdminModel returns a newsmodel for the mongo.
func NewAdminModel(config *config.Config) AdminModel {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, AdminCollectionName, config.CacheConf)
	return &customAdminModel{
		defaultAdminModel: newDefaultAdminModel(conn),
	}
}

var AdminSet = wire.NewSet(
	NewAdminModel,
)
