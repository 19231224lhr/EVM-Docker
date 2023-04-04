package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/vrf/internal/logic"
	"blockchain-crypto/crypto_api/vrf/internal/svc"
	"blockchain-crypto/crypto_api/vrf/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func vrfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VrfReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVrfLogic(r.Context(), svcCtx)
		resp, err := l.Vrf(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
