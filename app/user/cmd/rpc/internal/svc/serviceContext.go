package svc

import (
	"github.com/jinsoft/ainiok/app/identity/rpc/identity"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/internal/config"
	"github.com/jinsoft/ainiok/app/user/model"
	"github.com/tal-tech/go-zero/core/stores/redis"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	IdentityRpc identity.Identity
	// model
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.NewRedis(c.Redis.Host, c.Redis.Type),

		IdentityRpc: identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf)),

		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
