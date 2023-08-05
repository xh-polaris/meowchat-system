//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"

	"github.com/xh-polaris/meowchat-system/biz/adaptor"
)

func NewSystemServerImpl() (*adaptor.SystemServerImpl, error) {
	wire.Build(
		wire.Struct(new(adaptor.SystemServerImpl), "*"),
		AllProvider,
	)
	return nil, nil
}
