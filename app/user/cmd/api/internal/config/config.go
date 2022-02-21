package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	JwtAuth struct { // jwt鉴权配置
		AccessSecret string // jwt密钥
	}
	UserRpcConf zrpc.RpcClientConf
}
