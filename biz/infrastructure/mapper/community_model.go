package mapper

import (
	"context"
	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"meowchat-system/biz/infrastructure/config"
	"meowchat-system/biz/infrastructure/data/db"
)

const CommunityCollectionName = "community"

var _ CommunityModel = (*CustomCommunityModel)(nil)

type (
	// CommunityModel is an interface to be customized, add more methods here,
	// and implement the added methods in CustomCommunityModel.
	CommunityModel interface {
		communityModel
		ListCommunity(ctx context.Context, req *system.ListCommunityReq) ([]*db.Community, int64, error)
		DeleteCommunity(ctx context.Context, id string) error
	}

	CustomCommunityModel struct {
		*defaultCommunityModel
	}
)

func (c CustomCommunityModel) DeleteCommunity(ctx context.Context, id string) error {
	key := prefixCommunityCacheKey + id

	old := new(db.Community)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = c.conn.FindOneAndDelete(ctx, key, old, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	// delete children
	_, err = c.conn.DeleteMany(ctx, bson.M{
		"parentId": old.ID,
	})
	if err != nil {
		return err
	}

	// delete all cache
	err = c.conn.DelCache(ctx, prefixCommunityCacheKey+"*")
	if err != nil {
		return err
	}
	return nil
}

func (c CustomCommunityModel) ListCommunity(ctx context.Context, req *system.ListCommunityReq) ([]*db.Community, int64, error) {
	var resp []*db.Community

	filter := bson.M{}
	if req.ParentId != "" {
		pid, err := primitive.ObjectIDFromHex(req.ParentId)
		if err == nil {
			filter["parentId"] = pid
		}
	}

	findOptions := ToFindOptions(req.Page, req.PageSize, req.Sort)

	err := c.conn.Find(ctx, &resp, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}

	count, err := c.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return resp, count, nil
}

// NewCommunityModel returns a model for the mongo.
func NewCommunityModel(config *config.Config) CommunityModel {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CommunityCollectionName, config.CacheConf)
	return &CustomCommunityModel{
		defaultCommunityModel: newDefaultCommunityModel(conn),
	}
}

var CommunitySet = wire.NewSet(
	NewCommunityModel,
)
