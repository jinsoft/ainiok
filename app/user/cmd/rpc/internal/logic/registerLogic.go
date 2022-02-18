package logic

import (
	"context"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/pb"
	"github.com/jinsoft/ainiok/app/user/model"
	"github.com/jinsoft/ainiok/common/tool"
	"github.com/pkg/errors"

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

	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(errors.New("查询失败"), "mobile:%s, err:%v", in.Mobile, err)
	}

	if userInfo != nil {
		return nil, errors.Wrapf(errors.New("用户已经存在"), "用户已经存在 mobile:%s", in.Mobile)
	}

	user := new(model.User)
	user.Mobile = in.Mobile
	if len(user.Nickname) == 0 {
		user.Nickname = tool.RandomString(8)
	}
	if len(in.Password) > 0 {
		// 密码强度不够，需要升级
		user.Password = tool.MD5(in.Password)
	}
	if in.Gender < 1 || in.Gender > 3 {
		user.Gender = 3
	}

	insertResult, err := l.svcCtx.UserModel.Insert(user)
	if err != nil {
		return nil, errors.Wrapf(errors.New("查询失败"), "err:%v,user:%v", err, user)
	}
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(errors.New("获取插入的id失败"), "err:%v,user:%v", err, user)
	}

	// todo: 完成登陆

	return &pb.UserInfoReply{
		Id:       lastId,
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Gender:   tool.GetGenderInfo(user.Gender),
	}, nil
}
