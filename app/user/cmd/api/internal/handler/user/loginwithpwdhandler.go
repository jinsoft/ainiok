package user

import (
	"net/http"

	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/logic/user"
	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/svc"
	"github.com/jinsoft/ainiok/app/user/cmd/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginWithPwdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewLoginWithPwdLogic(r.Context(), svcCtx)
		resp, err := l.LoginWithPwd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
