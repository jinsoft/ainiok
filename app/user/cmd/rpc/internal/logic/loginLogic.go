package logic

import (
	"context"
	"fmt"
	"github.com/jinsoft/ainiok/app/identity/rpc/identity"
	"github.com/jinsoft/ainiok/app/user/model"
	"github.com/jinsoft/ainiok/common/tool"
	"github.com/pkg/errors"

	"github.com/jinsoft/ainiok/app/user/cmd/rpc/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/rpc/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginReply, error) {
	var (
		userId int64
		err    error
	)

	userId, err = l.LoginByPassword(in.Mobile, in.Password)
	if err != nil {
		return nil, err
	}

	// 生成登录token
	resp, err := l.svcCtx.IdentityRpc.GenerateToken(l.ctx, &identity.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(errors.New("生成token错误"), "IdentityRpc GenerateToken error, userId: %d", userId)
	}
	fmt.Println("******")
	fmt.Println(resp)

	loginRsp := &pb.LoginReply{
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}
	fmt.Println(loginRsp)
	return loginRsp, nil
	//return &pb.LoginReply{
	//	AccessToken:  resp.AccessToken,
	//	AccessExpire: resp.AccessExpire,
	//	RefreshAfter: resp.RefreshAfter,
	//}, nil
}

func (l *LoginLogic) LoginByPassword(mobile, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(errors.New("查询失败"), "账号密码登录查询失败, mobile:%s, err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(errors.New("用户不存在"), "mobile:%s", mobile)
	}
	if tool.MD5(password) != user.Password {
		return 0, errors.Wrapf(errors.New("密码错误"), "密码错误")
	}
	return user.Id, nil
}
