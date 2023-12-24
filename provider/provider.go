package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowchat-system/biz/application/service"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/admin"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/apply"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/community"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/news"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/notice"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/notification"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper/user_role"
)

var AllProvider = wire.NewSet(
	ApplicationSet,
	InfrastructureSet,
)

var ApplicationSet = wire.NewSet(
	service.SystemSet,
)

var InfrastructureSet = wire.NewSet(
	config.NewConfig,
	MapperSet,
)

var MapperSet = wire.NewSet(
	admin.AdminSet,
	apply.ApplySet,
	community.CommunitySet,
	news.NewsSet,
	notice.NoticeSet,
	user_role.UserRoleSet,
	notification.NotificationSet,
)
