package notification

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/pagination"
	"github.com/xh-polaris/gopkg/pagination/mongop"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/consts"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/util"
)

const (
	NotificationCollectionName = "notification"
)

var _ NotificationModel = (*customNotificationModel)(nil)

type (
	// NotificationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationModel.
	NotificationModel interface {
		notificationModel
		ListNotification(ctx context.Context, req *system.ListNotificationReq, sorter mongop.MongoCursor) ([]*db.Notification, int64, error)
		CleanNotification(ctx context.Context, userId string) error
		ReadNotification(ctx context.Context, notificationId string) error
		CountNotification(ctx context.Context, req *system.CountNotificationReq) (int64, error)
		ReadRange(ctx context.Context, req *system.ReadRangeNotificationReq) error
	}

	customNotificationModel struct {
		*defaultNotificationModel
	}
)

func (m customNotificationModel) ListNotification(ctx context.Context, req *system.ListNotificationReq, sorter mongop.MongoCursor) ([]*db.Notification, int64, error) {
	var data []*db.Notification
	popts := util.ParsePagination(req.GetPaginationOptions())
	p := mongop.NewMongoPaginator(pagination.NewRawStore(sorter), popts)
	f := &FilterOptions{
		OnlyUserId:     req.UserId,
		OnlyType:       req.Type,
		OnlyTargetType: req.TargetType,
	}
	filter := makeMongoFilter(f)
	sort, err := p.MakeSortOptions(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	if err = m.conn.Find(ctx, &data, filter, &options.FindOptions{
		Sort:  sort,
		Limit: req.PaginationOptions.Limit,
		Skip:  req.PaginationOptions.Offset,
	}); err != nil {
		return nil, 0, err
	}

	// 如果是反向查询，反转数据
	if *popts.Backward {
		for i := 0; i < len(data)/2; i++ {
			data[i], data[len(data)-i-1] = data[len(data)-i-1], data[i]
		}
	}
	if len(data) > 0 {
		err = p.StoreCursor(ctx, data[0], data[len(data)-1])
		if err != nil {
			return nil, 0, err
		}
	}

	count, err := m.conn.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return data, count, nil
}

func (m customNotificationModel) ReadRange(ctx context.Context, req *system.ReadRangeNotificationReq) error {
	lastOid, err := primitive.ObjectIDFromHex(req.GetLastId())
	if err != nil {
		return err
	}
	firstOid, err := primitive.ObjectIDFromHex(req.GetFirstId())
	if err != nil {
		return err
	}
	f := &FilterOptions{
		OnlyUserId:     req.UserId,
		OnlyType:       req.Type,
		OnlyTargetType: req.TargetType,
		FirstId:        &firstOid,
		LastId:         &lastOid,
	}
	filter := makeMongoFilter(f)

	if _, err := m.conn.UpdateManyNoCache(ctx, filter, bson.M{"$set": bson.M{consts.IsRead: true, consts.UpdateAt: time.Now()}}); err != nil {
		return err
	}
	return nil
}

func (m customNotificationModel) CleanNotification(ctx context.Context, userId string) error {

	filter := bson.M{
		consts.TargetUserId: userId,
		consts.IsRead:       bson.M{"$exists": false},
	}
	_, err := m.conn.UpdateManyNoCache(ctx, filter, bson.M{"$set": bson.M{consts.IsRead: true, consts.UpdateAt: time.Now()}})
	return err
}

func (m customNotificationModel) ReadNotification(ctx context.Context, notificationId string) error {
	oid, err := primitive.ObjectIDFromHex(notificationId)
	if err != nil {
		return mapper.ErrInvalidObjectId
	}

	key := prefixNotificationCacheKey + notificationId
	_, err = m.conn.UpdateByID(ctx, key, oid, bson.M{"$set": bson.M{consts.IsRead: true, consts.UpdateAt: time.Now()}})
	return err
}

func (m customNotificationModel) CountNotification(ctx context.Context, req *system.CountNotificationReq) (int64, error) {
	isRead := false
	f := &FilterOptions{
		OnlyUserId:     &req.UserId,
		OnlyType:       req.Type,
		OnlyTargetType: req.TargetType,
		IsRead:         &isRead,
	}
	filter := makeMongoFilter(f)
	return m.conn.CountDocuments(ctx, filter)
}

// NewNotificationModel returns a Notification-model for the mongo.
func NewNotificationModel(config *config.Config) NotificationModel {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, NotificationCollectionName, config.CacheConf)
	return &customNotificationModel{
		defaultNotificationModel: newDefaultNotificationModel(conn),
	}
}

var NotificationSet = wire.NewSet(
	NewNotificationModel,
)
