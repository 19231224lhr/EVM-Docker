package handler

import (
	"net/http"

	"blockchain-crypto/crypto_api/key_exchange/internal/logic"
	"blockchain-crypto/crypto_api/key_exchange/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateKeyEcdhHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGenerateKeyEcdhLogic(r.Context(), svcCtx)
		resp, err := l.GenerateKeyEcdh()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
