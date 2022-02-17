package logic

import (
	"context"

	"github.com/jinsoft/ainiok/app/user/cmd/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.UserInfoReply, error) {
	// todo: add your logic here and delete this line

	return &pb.UserInfoReply{
		Id:       111,
		Nickname: "ainiok",
		Mobile:   "18888888888",
		Gender:   "保密",
	}, nil
}
