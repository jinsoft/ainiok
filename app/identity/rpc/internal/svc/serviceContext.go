package svc

import (
	"github.com/jinsoft/ainiok/app/identity/rpc/internal/config"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.NewRedis(c.Redis.Host, c.Redis.Type),
	}
}
