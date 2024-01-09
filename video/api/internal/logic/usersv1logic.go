package logic

import (
	"context"

	"go-zero-test/video/api/internal/svc"
	"go-zero-test/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Usersv1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersv1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Usersv1Logic {
	return &Usersv1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Usersv1Logic) Usersv1() (resp []types.UserV1, err error) {
	// todo: add your logic here and delete this line

	return
}
