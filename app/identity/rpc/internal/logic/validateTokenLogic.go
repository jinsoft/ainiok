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

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {

	tokenKey := fmt.Sprintf(globalKey.CacheUserTokenKey, in.UserId)
	cacheToken, err := l.svcCtx.RedisClient.Get(tokenKey)
	if err != nil {
		return nil, errors.Wrapf(errors.New("redis 获取token失败"), "userId:%d, err:%v", in.UserId, err)
	}
	if cacheToken != in.Token {
		return nil, errors.Wrapf(errors.New("token 无效"), "token 无效")
	}
	return &pb.ValidateTokenResp{
		Ok: true,
	}, nil
}
