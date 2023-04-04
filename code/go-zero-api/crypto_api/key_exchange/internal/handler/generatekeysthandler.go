package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/key_exchange/internal/logic"
	"blockchain-crypto/crypto_api/key_exchange/internal/svc"
	"blockchain-crypto/crypto_api/key_exchange/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateKeyStHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateKeyStReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGenerateKeyStLogic(r.Context(), svcCtx)
		resp, err := l.GenerateKeySt(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
