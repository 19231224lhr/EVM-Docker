package logic

import (
	"blockchain-crypto/vdf/wesolowski_rust"
	"context"
	"fmt"

	"blockchain-crypto/crypto_api/vdf/internal/svc"
	"blockchain-crypto/crypto_api/vdf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Vdf_wrustLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVdf_wrustLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Vdf_wrustLogic {
	return &Vdf_wrustLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Vdf_wrustLogic) Vdf_wrust(req *types.WrustReq) (resp *types.WrustRes, err error) {
	// todo: add your logic here and delete this line
	switch req.Name {
	case "1":
		out := wesolowski_rust.Execute(req.Challenge, req.Iterations)
		fmt.Println(out)
		fmt.Println("out_string", string(out))
		return &types.WrustRes{
			Out: string(out),
		}, nil
	case "2":
		out2 := wesolowski_rust.Verify(req.Challenge, req.Iterations, req.Proof)
		fmt.Println(out2)
		return &types.WrustRes{
			Out2: out2,
		}, nil
	default:
		fmt.Println("wrong num")
		return nil, nil
	}
}
