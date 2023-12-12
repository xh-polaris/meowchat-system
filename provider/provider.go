package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowchat-system/biz/application/service"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/mapper"
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
	mapper.AdminSet,
	mapper.ApplySet,
	mapper.CommunitySet,
	mapper.NewsSet,
	mapper.NoticeSet,
	mapper.UserRoleSet,
	mapper.NotificationSet,
)
