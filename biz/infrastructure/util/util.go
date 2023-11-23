package util

import (
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
)

func ConvertAdmin(in *db.Admin) *system.Admin {
	return &system.Admin{
		Id:          in.ID.Hex(),
		CommunityId: in.CommunityId,
		Name:        in.Name,
		Title:       in.Title,
		Phone:       in.Phone,
		Wechat:      in.Wechat,
		AvatarUrl:   in.AvatarUrl,
	}
}

func ConvertNotice(in *db.Notice) *system.Notice {
	return &system.Notice{
		Id:          in.ID.Hex(),
		CommunityId: in.CommunityId,
		Text:        in.Text,
		CreateAt:    in.CreateAt.Unix(),
		UpdateAt:    in.UpdateAt.Unix(),
	}
}

func ConvertNews(in *db.News) *system.News {
	return &system.News{
		Id:          in.ID.Hex(),
		CommunityId: in.CommunityId,
		ImageUrl:    in.ImageUrl,
		LinkUrl:     in.LinkUrl,
		Type:        in.Type,
		IsPublic:    in.IsPublic,
	}
}

func ConvertCommunity(in *db.Community) *system.Community {
	pid := ""
	if in.ParentId != primitive.NilObjectID {
		pid = in.ParentId.Hex()
	}

	return &system.Community{
		Id:       in.ID.Hex(),
		Name:     in.Name,
		ParentId: pid,
	}
}

func ConvertNotification(in *db.Notification) *system.Notification {
	return &system.Notification{
		NotificationId:  in.NotificationId.Hex(),
		TargetUserId:    in.TargetUserId,
		SourceUserId:    in.SourceUserId,
		SourceContentId: in.SourceContentId,
		Type:            in.Type,
		Text:            in.Text,
		CreateAt:        in.CreateAt.Unix(),
		IsRead:          in.IsRead,
	}
}

func ConvertNotifications(in []*db.Notification) []*system.Notification {
	res := make([]*system.Notification, 0)
	for _, v := range in {
		res = append(res, &system.Notification{
			NotificationId:  v.NotificationId.Hex(),
			TargetUserId:    v.TargetUserId,
			SourceUserId:    v.SourceUserId,
			SourceContentId: v.SourceContentId,
			Type:            v.Type,
			Text:            v.Text,
			CreateAt:        v.CreateAt.Unix(),
			IsRead:          v.IsRead,
		})
	}
	return res
}
