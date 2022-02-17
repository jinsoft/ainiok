package user

import (
	"context"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/user"
	"github.com/pkg/errors"

	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	registerResp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
		Nickname: "ainiok",
		Password: "123456",
		Mobile:   "1888888888",
		Gender:   1,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req:%v", req)
	}

	return &types.RegisterResp{
		Id:       registerResp.Id,
		Nickname: registerResp.Nickname,
		Mobile:   registerResp.Mobile,
		Gender:   1,
	}, nil
}
