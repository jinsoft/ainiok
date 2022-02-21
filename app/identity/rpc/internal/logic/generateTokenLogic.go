package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jinsoft/ainiok/common/globalKey"
	"github.com/pkg/errors"
	"time"

	"github.com/jinsoft/ainiok/app/identity/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/identity/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(errors.New("生成token失败"), "getJwtToken err, userId:%d", in.UserId)
	}

	// 存入redis
	tokenKey := fmt.Sprintf(globalKey.CacheUserTokenKey, in.UserId)
	err = l.svcCtx.RedisClient.Setex(tokenKey, accessToken, int(accessExpire))
	if err != nil {
		return nil, errors.Wrapf(errors.New("redis设置token失败"), "SetnxEx err userId:%d, err:%v", in.UserId, err)
	}
	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["jwtUserId"] = userId

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
