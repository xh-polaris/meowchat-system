package mapper

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
)

const (
	NotificationId             = "_id"
	CreateAt                   = "createAt"
	TargetUserId               = "targetUserId"
	IsRead                     = "isRead"
	UpdateAt                   = "updateAt"
	NotificationCollectionName = "notification"
)

var _ NotificationModel = (*customNotificationModel)(nil)

type (
	// NotificationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationModel.
	NotificationModel interface {
		notificationModel
		ListNotification(ctx context.Context, req *system.ListNotificationReq) ([]*db.Notification, int64, error)
		CleanNotification(ctx context.Context, userId string) error
		ReadNotification(ctx context.Context, notificationId string) error
		CountNotification(ctx context.Context, userId string) (int64, error)
	}

	customNotificationModel struct {
		*defaultNotificationModel
	}
)

func (m customNotificationModel) ListNotification(ctx context.Context, req *system.ListNotificationReq) ([]*db.Notification, int64, error) {
	var resp []*db.Notification

	filter := bson.M{
		"userId": req.UserId,
	}

	sort := bson.E{Key: CreateAt, Value: -1}

	findOptions := &options.FindOptions{
		Limit: req.PaginationOptions.Limit,
		Skip:  req.PaginationOptions.Offset,
		Sort:  sort,
	}

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

func (m customNotificationModel) CleanNotification(ctx context.Context, userId string) error {

	filter := bson.M{
		TargetUserId: userId,
		IsRead:       false,
	}
	_, err := m.conn.UpdateManyNoCache(ctx, filter, bson.M{"$set": bson.M{IsRead: true, UpdateAt: time.Now()}})
	return err
}

func (m customNotificationModel) ReadNotification(ctx context.Context, notificationId string) error {
	key := prefixNotificationCacheKey + notificationId
	_, err := m.conn.UpdateByID(ctx, key, notificationId, bson.M{"$set": bson.M{IsRead: true, UpdateAt: time.Now()}})
	return err
}

func (m customNotificationModel) CountNotification(ctx context.Context, userId string) (int64, error) {
	filter := bson.M{
		TargetUserId: userId,
		IsRead:       false,
	}
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
