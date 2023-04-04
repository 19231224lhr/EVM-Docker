package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/secret_sharing/internal/logic"
	"blockchain-crypto/crypto_api/secret_sharing/internal/svc"
	"blockchain-crypto/crypto_api/secret_sharing/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SecretSSReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSSLogic(r.Context(), svcCtx)
		resp, err := l.SS(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
