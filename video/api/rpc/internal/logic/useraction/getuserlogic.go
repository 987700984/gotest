package useractionlogic

import (
	"context"

	"go-zero-test/video/api/rpc/internal/svc"
	"go-zero-test/video/api/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserInfoReponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfoReponse{}, nil
}
