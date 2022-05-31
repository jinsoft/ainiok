package logic

import (
	"context"
	"github.com/jinsoft/ainiok/app/user/model"
	"github.com/jinsoft/ainiok/common/tool"
	"github.com/jinsoft/ainiok/common/xerr"
	"github.com/pkg/errors"

	"github.com/jinsoft/ainiok/app/user/cmd/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var UserNotExistsError = xerr.NewErrCode(xerr.UserNotExists)

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.IdReq) (*pb.UserInfoReply, error) {

	user, err := l.svcCtx.UserModel.FindOne(in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errors.New("数据库错误"), "查询失败, id:%d, err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(UserNotExistsError, "id:%d", in.Id)
	}

	return &pb.UserInfoReply{
		Id:       user.Id,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Gender:   tool.GetGenderInfo(user.Gender),
	}, nil
}
