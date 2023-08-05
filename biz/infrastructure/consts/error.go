package consts

import (
	"google.golang.org/grpc/status"
	"meowchat-system/biz/infrastructure/mapper"
)

var (
	ErrNotFound                 = status.Error(10001, "no such element")
	ErrInvalidObjectId          = status.Error(10002, "invalid objectId")
	ErrCommunityIdNotFound      = status.Error(10003, "communityId not found")
	ErrChildCommunityNotAllowed = status.Error(10004, "child community not allowed")
)

func Switch(err error) error {
	switch err {
	case mapper.ErrNotFound:
		return ErrNotFound
	case mapper.ErrInvalidObjectId:
		return ErrInvalidObjectId
	default:
		return err
	}
}
