package user

import (
	"context"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/user"
	"github.com/jinsoft/ainiok/common/ctxdata"

	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	infoResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		Id:       userId,
		Nickname: infoResp.Nickname,
		Mobile:   infoResp.Mobile,
		Gender:   infoResp.Gender,
	}, nil
}
