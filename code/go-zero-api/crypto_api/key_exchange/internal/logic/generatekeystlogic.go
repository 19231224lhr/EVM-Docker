package logic

import (
	"blockchain-crypto/crypto_api/key_exchange/internal/svc"
	"blockchain-crypto/crypto_api/key_exchange/internal/types"
	"blockchain-crypto/key_exchange/stealth_address"
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateKeyStLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateKeyStLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateKeyStLogic {
	return &GenerateKeyStLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func publicKeyToJson(pubKey stealth_address.PublicKey) (string, error) {
	jsonBytes, err := json.Marshal(pubKey)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func jsonToPublicKey(jsonString string) (stealth_address.PublicKey, error) {
	var decodedPubKey stealth_address.PublicKey
	jsonBytes := []byte(jsonString)
	err := json.Unmarshal(jsonBytes, &decodedPubKey)
	if err != nil {
		return stealth_address.PublicKey{}, err
	}
	return decodedPubKey, nil
}

func (l *GenerateKeyStLogic) GenerateKeySt(req *types.GenerateKeyStReq) (resp *types.GenerateKeyStRes, err error) {
	// todo: add your logic here and delete this line
	switch req.Name {
	case "1":
		pri1, pub1, pri2, pub2 := stealth_address.RecCalculateKeyPairs(req.Str1, req.Str2)

		jsonStr1, _ := publicKeyToJson(*pub1)

		jsonStr2, _ := publicKeyToJson(*pub2)

		return &types.GenerateKeyStRes{
			Priv1: pri1.Text(10),
			Pub1:  jsonStr1,
			Priv2: pri2.Text(10),
			Pub2:  jsonStr2,
		}, nil
	case "2":
		r, _ := new(big.Int).SetString(req.Req_r, 10)
		pub1 := req.Req_pub1
		// 将 JSON 字符串转换回 PublicKey 结构体
		var pub1_decoded stealth_address.PublicKey

		pub1_decoded, _ = jsonToPublicKey(pub1)
		pub2 := req.Req_pub2
		// 将 JSON 字符串转换回 PublicKey 结构体
		var pub2_decoded stealth_address.PublicKey

		pub2_decoded, _ = jsonToPublicKey(pub2)
		// 使用defer在panic时捕获异常
		iserror := true
		defer func() {
			if err := recover(); err != nil {
				iserror = false
			}
		}()
		P := stealth_address.SendCalculateObfuscateAddress(r, &pub1_decoded, &pub2_decoded)
		R := stealth_address.SendCalculatePublicKey(r)
		Pjson, _ := publicKeyToJson(*P)
		Rjson, _ := publicKeyToJson(*R)

		if iserror == false {
			return &types.GenerateKeyStRes{
				P: "error point",
				R: "error point",
			}, nil
		}
		fmt.Println(r, pub1_decoded, pub2_decoded, R)
		return &types.GenerateKeyStRes{
			P: Pjson,
			R: Rjson,
		}, nil

	case "3":
		P_encoded, _ := jsonToPublicKey(req.Req_P)
		R_encoded, _ := jsonToPublicKey(req.Req_R)
		// string -> private key(big int)
		var num1 big.Int
		num1.SetString(req.Req_priv1, 10)
		// string json -> public key
		var num2 stealth_address.PublicKey
		num2, _ = jsonToPublicKey(req.Req_pub2)
		var num3 big.Int
		num3.SetString(req.Req_priv2, 10)
		t := stealth_address.RecCalculateObfuscateAddress(&P_encoded, &R_encoded, &num1, &num2)
		fmt.Println(t)
		priv := stealth_address.RecCalculateAddressPrivatekey(&num1, &num3, &R_encoded)
		fmt.Println(priv)

		return &types.GenerateKeyStRes{
			T:    t,
			Priv: priv.Text(10),
		}, nil
	default:
		fmt.Println("wrong num")
		return nil, nil
	}
}
