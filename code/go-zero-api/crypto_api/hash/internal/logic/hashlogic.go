package logic

import (
	"context"
	"fmt"

	"blockchain-crypto/hash"

	"blockchain-crypto/crypto_api/hash/internal/svc"
	"blockchain-crypto/crypto_api/hash/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HashLogic {
	return &HashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HashLogic) Hash(req *types.HashReq) (resp *types.HashRes, err error) {

	input := []byte(req.HashMessage)
	output, err := hash.Hash(req.HashType, input)

	if err != nil {
		fmt.Println(req.HashType, err)
		return nil, err
	}

	return &types.HashRes{
		Res: output,
	}, nil
}
