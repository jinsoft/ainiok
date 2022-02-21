package verify

import (
	"context"

	"github.com/jinsoft/ainiok/app/identity/api/internal/svc"
	"github.com/jinsoft/ainiok/app/identity/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type TokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) TokenLogic {
	return TokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TokenLogic) Token(req types.VerifyTokenReq) (resp *types.VerifyTokenResp, err error) {
	// todo: add your logic here and delete this line

	return
}
