package provider

import (
	"github.com/google/wire"
	"meowchat-system/biz/application/service"
	"meowchat-system/biz/infrastructure/config"
	"meowchat-system/biz/infrastructure/mapper"
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
)
