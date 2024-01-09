// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package useraction

import (
	"context"

	"go-zero-test/video/api/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest       = user.IdRequest
	UserInfoReponse = user.UserInfoReponse

	UserAction interface {
		GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoReponse, error)
		GetAAF(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoReponse, error)
		GetUser1(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoReponse, error)
	}

	defaultUserAction struct {
		cli zrpc.Client
	}
)

func NewUserAction(cli zrpc.Client) UserAction {
	return &defaultUserAction{
		cli: cli,
	}
}

func (m *defaultUserAction) GetUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoReponse, error) {
	client := user.NewUserActionClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUserAction) GetAAF(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoReponse, error) {
	client := user.NewUserActionClient(m.cli.Conn())
	return client.GetAAF(ctx, in, opts...)
}

func (m *defaultUserAction) GetUser1(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*UserInfoReponse, error) {
	client := user.NewUserActionClient(m.cli.Conn())
	return client.GetUser1(ctx, in, opts...)
}
