package mq

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	consumer2 "github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper"
)

type Notification struct {
	UserId   string
	ParentId string
	Type     string
	Text     string
	CreateAt int64
}

var (
	NotificationModel mapper.NotificationModel
	umu               sync.Mutex
)

func NewMqConsumer(c *config.Config) rocketmq.PushConsumer {
	consumer, err := rocketmq.NewPushConsumer(
		consumer2.WithNsResolver(primitive.NewPassthroughResolver(c.RocketMq.URL)),
		consumer2.WithGroupName(c.RocketMq.GroupName),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = consumer.Subscribe(
		"notification",
		consumer2.MessageSelector{},
		func(ctx context.Context, ext ...*primitive.MessageExt) (consumer2.ConsumeResult, error) {
			for i := range ext {
				err := NotificationMessageHandler(c, ext[i].Body)
				if err != nil {
					logx.Alert(fmt.Sprintf("%v", err))
				}
			}
			return consumer2.ConsumeSuccess, nil
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	err = consumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	return consumer
}

func NotificationMessageHandler(c *config.Config, body []byte) interface{} {
	checkSingletonModel(c)
	msg := new(system.Notification)
	err := jsonx.Unmarshal(body, &msg)
	if err != nil {
		return err
	}
	notification := &db.Notification{
		TargetUserId:    msg.GetTargetUserId(),
		SourceUserId:    msg.GetSourceUserId(),
		SourceContentId: msg.GetSourceContentId(),
		Type:            msg.GetType(),
		Text:            msg.GetText(),
		IsRead:          msg.GetIsRead(),
		CreateAt:        time.Now(),
		UpdateAt:        time.Now(),
	}
	err = NotificationModel.Insert(context.Background(), notification)
	if err != nil {
		return err
	}
	return nil
}

func checkSingletonModel(c *config.Config) {
	if NotificationModel == nil {
		umu.Lock()
		if NotificationModel == nil {
			Model := mapper.NewNotificationModel(c)
			NotificationModel = Model
		}
		umu.Unlock()
	}
}
