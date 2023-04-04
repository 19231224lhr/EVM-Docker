package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/vdf/internal/logic"
	"blockchain-crypto/crypto_api/vdf/internal/svc"
	"blockchain-crypto/crypto_api/vdf/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func vdf_pietrzakHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PietrzakReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVdf_pietrzakLogic(r.Context(), svcCtx)
		resp, err := l.Vdf_pietrzak(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
