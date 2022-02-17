package svc

import (
	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/config"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/user"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
