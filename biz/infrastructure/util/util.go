package util

import (
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
