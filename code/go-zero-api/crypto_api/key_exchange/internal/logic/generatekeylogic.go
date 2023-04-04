package logic

import (
	"blockchain-crypto/key_exchange/dh"
	"context"
	"fmt"

	"blockchain-crypto/crypto_api/key_exchange/internal/svc"
	"blockchain-crypto/crypto_api/key_exchange/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateKeyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateKeyLogic {
	return &GenerateKeyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateKeyLogic) GenerateKey() (resp *types.GenerateKeyRes, err error) {
	// todo: add your logic here and delete this line
	// 1 {string. string} -> group
	// 使用2048位群
	group := dh.Init_4096()

	// 2 res = function(group)
	// Alice创建公私钥对
	alicePrivate, alicePublic, err := group.GenerateKey()
	if err != nil {
		fmt.Printf("Failed to generate alice's private / public key pair: %s", err)
	}

	// Bob创建公私钥对
	bobPrivate, bobPublic, err := group.GenerateKey()
	if err != nil {
		fmt.Printf("Failed to generate bob's private / public key pair: %s", err)
	}

	//Alice计算会话密钥
	secretAlice := group.ComputeSecret(alicePrivate, bobPublic)
	//Bob计算会话密钥
	secretBob := group.ComputeSecret(bobPrivate, alicePublic)
	fmt.Println("============================================")
	return &types.GenerateKeyRes{
		SecretA: secretAlice.Text(10), // turn to string
		SecretB: secretBob.Text(10),   // turn to string
	}, nil
}
