package logic

import (
	"context"

	"github.com/jinsoft/ainiok/app/user/cmd/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.IdReq) (*pb.UserInfoReply, error) {
	// todo: add your logic here and delete this line

	return &pb.UserInfoReply{}, nil
}
