package db

import (
	"time"

	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	NotificationId  primitive.ObjectID            `bson:"_id,omitempty" json:"_id,omitempty"`
	TargetUserId    string                        `bson:"targetUserId,omitempty" json:"targetUserId,omitempty"`
	SourceUserId    string                        `bson:"sourceUserId,omitempty" json:"sourceUserId,omitempty"`
	SourceContentId string                        `bson:"sourceContentId,omitempty" json:"sourceContentId,omitempty"`
	Type            system.NotificationType       `bson:"type,omitempty" json:"type,omitempty"`
	TargetType      system.NotificationTargetType `bson:"targetType,omitempty" json:"targetType,omitempty"`
	Text            string                        `bson:"text,omitempty" json:"text,omitempty"`
	CreateAt        time.Time                     `bson:"createAt,omitempty" json:"createAt,omitempty"`
	UpdateAt        time.Time                     `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	IsRead          bool                          `bson:"isRead,omitempty" json:"isRead,omitempty"`
}
