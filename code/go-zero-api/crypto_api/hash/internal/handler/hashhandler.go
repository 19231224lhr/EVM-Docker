package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/hash/internal/logic"
	"blockchain-crypto/crypto_api/hash/internal/svc"
	"blockchain-crypto/crypto_api/hash/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func hashHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HashReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHashLogic(r.Context(), svcCtx)
		resp, err := l.Hash(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
