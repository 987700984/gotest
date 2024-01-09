package handler

import (
	"net/http"

	"github.com/987700984/gotest/video/api/internal/logic"
	"github.com/987700984/gotest/video/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func usersv2Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUsersv2Logic(r.Context(), svcCtx)
		resp, err := l.Usersv2()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
