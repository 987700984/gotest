// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero-test/user/rpc/internal/logic/user"
	"go-zero-test/user/rpc/internal/svc"
	"go-zero-test/user/rpc/types/user"
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
