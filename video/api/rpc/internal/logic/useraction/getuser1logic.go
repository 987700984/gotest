package useractionlogic

import (
	"context"

	"github.com/987700984/gotest/video/api/rpc/internal/svc"
	"github.com/987700984/gotest/video/api/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUser1Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUser1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUser1Logic {
	return &GetUser1Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUser1Logic) GetUser1(in *user.IdRequest) (*user.UserInfoReponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoReponse{}, nil
}
