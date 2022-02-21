package verify

import (
	"net/http"

	"github.com/jinsoft/ainiok/app/identity/api/internal/logic/verify"
	"github.com/jinsoft/ainiok/app/identity/api/internal/svc"
	"github.com/jinsoft/ainiok/app/identity/api/internal/types"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func TokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := verify.NewTokenLogic(r.Context(), svcCtx)
		resp, err := l.Token(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
