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
	"time"
)

const NewsCollectionName = "news"

var _ NewsModel = (*customNewsModel)(nil)

type (
	// NewsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNewsModel.
	NewsModel interface {
		newsModel
		UpdateNews(ctx context.Context, req *system.UpdateNewsReq) error
		ListNews(ctx context.Context, req *system.ListNewsReq) ([]*db.News, int64, error)
	}

	customNewsModel struct {
		*defaultNewsModel
	}
)

func (m customNewsModel) ListNews(ctx context.Context, req *system.ListNewsReq) ([]*db.News, int64, error) {
	var resp []*db.News

	filter := bson.M{
		"$or": []bson.M{
			{"communityId": req.CommunityId},
			{"isPublic": 1},
		},
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

func (m customNewsModel) UpdateNews(ctx context.Context, req *system.UpdateNewsReq) error {
	key := prefixNewsCacheKey + req.Id

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return ErrInvalidObjectId
	}

	filter := bson.M{
		"_id": oid,
	}
	set := bson.M{
		"type":     req.Type,
		"imageUrl": req.ImageUrl,
		"linkUrl":  req.LinkUrl,
		"updateAt": time.Now(),
		"isPublic": req.IsPublic,
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

// NewNewsModel returns a newsmodel for the mongo.
func NewNewsModel(config *config.Config) NewsModel {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, NewsCollectionName, config.CacheConf)
	return &customNewsModel{
		defaultNewsModel: newDefaultNewsModel(conn),
	}
}

var NewsSet = wire.NewSet(
	NewNewsModel,
)
