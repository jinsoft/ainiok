package logic

import (
	"context"
	"fmt"
	"github.com/jinsoft/ainiok/common/globalKey"
	"github.com/pkg/errors"

	"github.com/jinsoft/ainiok/app/identity/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/identity/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClearTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearTokenLogic {
	return &ClearTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClearTokenLogic) ClearToken(in *pb.ClearTokenReq) (*pb.ClearTokenResp, error) {

	tokenKey := fmt.Sprintf(globalKey.CacheUserTokenKey, in.UserId)
	if _, err := l.svcCtx.RedisClient.Del(tokenKey); err != nil {
		return nil, errors.Wrapf(errors.New("删除token失败"), "userId:%d, err:%v", in.UserId, err)
	}
	return &pb.ClearTokenResp{
		Ok: true,
	}, nil
}
