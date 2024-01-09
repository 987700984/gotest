package userlogic

import (
	"context"

	"go-zero-test/user/rpc/internal/svc"
	"go-zero-test/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginReponse, error) {
	// todo: add your logic here and delete this line

	return &user.LoginReponse{Email: "豆腐干豆腐干"}, nil
}
