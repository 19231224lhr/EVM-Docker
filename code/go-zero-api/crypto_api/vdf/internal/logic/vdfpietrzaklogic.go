package logic

import (
	"blockchain-crypto/vdf/pietrzak"
	"context"
	"fmt"

	"blockchain-crypto/crypto_api/vdf/internal/svc"
	"blockchain-crypto/crypto_api/vdf/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Vdf_pietrzakLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVdf_pietrzakLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Vdf_pietrzakLogic {
	return &Vdf_pietrzakLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Vdf_pietrzakLogic) Vdf_pietrzak(req *types.PietrzakReq) (resp *types.PietrzakRes, err error) {
	// todo: add your logic here and delete this line
	switch req.Name {
	case "1":
		out := pietrzak.Execute(req.Challenge, req.Iterations)
		fmt.Println(out)
		fmt.Println("out_string", string(out))
		return &types.PietrzakRes{
			Out: string(out),
		}, nil
	case "2":
		out2 := pietrzak.Verify(req.Challenge, req.Iterations, req.Proof)
		fmt.Println(out2)
		return &types.PietrzakRes{
			Out2: out2,
		}, nil
	default:
		fmt.Println("wrong num")
		return nil, nil
	}
}
