package rpcServer

import (
	"context"
	"github.com/jinsoft/ainiok/common/xerr"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)

	if err != nil {
		causeErr := errors.Cause(err)
		logx.WithContext(ctx).Errorf("[RPC-SRV-ERR] %+v", err)
		if e, ok := causeErr.(*xerr.CodeError); ok {
			// 自定义错误类型
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		}
	}
	return resp, err
}
