package logic

import (
	"context"

	"github.com/987700984/gotest/video/api/internal/svc"
	"github.com/987700984/gotest/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Usersv2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersv2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Usersv2Logic {
	return &Usersv2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Usersv2Logic) Usersv2() (resp []types.UserV2, err error) {
	// todo: add your logic here and delete this line

	return
}
