// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/987700984/gotest/user/rpc/internal/logic/user"
	"github.com/987700984/gotest/user/rpc/internal/svc"
	"github.com/987700984/gotest/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginReponse, error) {
	l := userlogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
