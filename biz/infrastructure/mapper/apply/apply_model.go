package apply

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/system"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/xh-polaris/meowchat-system/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-system/biz/infrastructure/data/db"
)

const ApplyCollectionName = "apply"

var _ ApplyModel = (*customApplyModel)(nil)

type (
	// ApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplyModel.
	ApplyModel interface {
		applyModel
		FindAllApplyByCommunityId(ctx context.Context, req *system.ListApplyReq) ([]*db.Apply, error)
	}

	customApplyModel struct {
		*defaultApplyModel
	}
)

func (m customApplyModel) FindAllApplyByCommunityId(ctx context.Context, req *system.ListApplyReq) ([]*db.Apply, error) {
	var resp []*db.Apply
	filter := bson.M{
		"communityId": req.CommunityId,
	}
	if err := m.conn.Find(ctx, &resp, filter); err != nil {
		return nil, err
	}
	return resp, nil
}

// NewApplyModel returns a model for the mongo.
func NewApplyModel(config *config.Config) ApplyModel {
	conn := mon.MustNewModel(config.Mongo.URL, config.Mongo.DB, ApplyCollectionName)
	return &customApplyModel{
		defaultApplyModel: newDefaultApplyModel(conn),
	}
}

var ApplySet = wire.NewSet(
	NewApplyModel,
)
