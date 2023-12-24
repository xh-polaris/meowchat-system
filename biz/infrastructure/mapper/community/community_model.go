package community

import (
	"context"
	"errors"
	"time"

	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper"
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
		InsertRoot(ctx context.Context, data *db.Community) error
	}

	CustomCommunityModel struct {
		*defaultCommunityModel
	}
)

func (m *CustomCommunityModel) InsertRoot(ctx context.Context, data *db.Community) error {
	if !data.ParentId.IsZero() {
		return errors.New("not root community")
	}
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	s, err := m.conn.StartSession()
	if err != nil {
		return err
	}
	_, err = s.WithTransaction(ctx, func(ctx mongo.SessionContext) (interface{}, error) {
		_, err = m.conn.InsertOne(ctx, prefixCommunityCacheKey+data.ID.Hex(), data)
		if err != nil {
			return nil, err
		}
		data.ParentId = data.ID
		data.ID = primitive.NewObjectID()
		_, err = m.conn.InsertOne(ctx, prefixCommunityCacheKey+data.ID.Hex(), data)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *CustomCommunityModel) DeleteCommunity(ctx context.Context, id string) error {
	key := prefixCommunityCacheKey + id

	old := new(db.Community)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = m.conn.FindOneAndDelete(ctx, key, old, bson.M{"_id": oid})
	if err != nil {
		return err
	}

	// delete children
	_, err = m.conn.DeleteMany(ctx, bson.M{
		"parentId": old.ID,
	})
	if err != nil {
		return err
	}

	// delete all cache
	err = m.conn.DelCache(ctx, prefixCommunityCacheKey+"*")
	if err != nil {
		return err
	}
	return nil
}

func (m *CustomCommunityModel) ListCommunity(ctx context.Context, req *system.ListCommunityReq) ([]*db.Community, int64, error) {
	var resp []*db.Community

	filter := bson.M{}
	if req.ParentId != "" {
		pid, err := primitive.ObjectIDFromHex(req.ParentId)
		if err == nil {
			filter["parentId"] = pid
		}
	}

	findOptions := mapper.ToFindOptions(req.Page, req.PageSize, req.Sort)

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
