package logic

import (
	"blockchain-crypto/crypto_api/key_exchange/internal/svc"
	"blockchain-crypto/crypto_api/key_exchange/internal/types"
	"blockchain-crypto/key_exchange/ecdh"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateKeyEcdhLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateKeyEcdhLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateKeyEcdhLogic {
	return &GenerateKeyEcdhLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateKeyEcdhLogic) GenerateKeyEcdh() (resp *types.GenerateKeyEcdh, err error) {
	// todo: add your logic here and delete this line
	//需要密钥协商的双方各自计算公私钥对（privatekey1, publickey1）,(privatekey2, publickey2),其中公钥是椭圆曲线上由G和私钥生成的点的坐标
	privatekey1, publickey1 := ecdh.CalculateKeypair()
	privatekey2, publickey2 := ecdh.CalculateKeypair()

	//需要密钥协商的双方获得对方的公钥后并利用自身的私钥计算协商后的密钥
	shared1 := ecdh.CalculateNegotiationKey(publickey2, privatekey1)
	shared2 := ecdh.CalculateNegotiationKey(publickey1, privatekey2)

	// fmt.Printf("%x\n", shared1)
	// fmt.Printf("%x\n", shared2)

	return &types.GenerateKeyEcdh{
		SecretA: string(shared1[:]),
		SecretB: string(shared2[:]),
	}, nil
}
