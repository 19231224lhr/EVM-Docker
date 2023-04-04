package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/secret_sharing/internal/logic"
	"blockchain-crypto/crypto_api/secret_sharing/internal/svc"
	"blockchain-crypto/crypto_api/secret_sharing/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PVSSHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SecretPvssReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPVSSLogic(r.Context(), svcCtx)
		resp, err := l.PVSS(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
