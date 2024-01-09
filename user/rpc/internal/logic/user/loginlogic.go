package userlogic

import (
	"context"

	"github.com/987700984/gotest/user/rpc/internal/svc"
	"github.com/987700984/gotest/user/rpc/types/user"

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
