package logic

import (
	"blockchain-crypto/crypto_api/secret_sharing/internal/svc"
	"blockchain-crypto/crypto_api/secret_sharing/internal/types"
	"blockchain-crypto/secret_sharing/ss"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type SSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SSLogic {
	return &SSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SSLogic) SS(req *types.SecretSSReq) (resp *types.SecretSSRes, err error) {
	// todo: add your logic here and delete this line
	switch req.Name {
	case "1":
		n := req.N
		k := req.K
		secret1 := req.Secret // string
		// string -> []byte
		secret, err := base64.StdEncoding.DecodeString(secret1)
		fmt.Println("==================secret", secret)
		shares := ss.Split(n, k, secret)
		fmt.Println("==================shares", shares)
		// 将 []Share 转换为 []byte
		bytesShares := make([][]byte, len(shares))
		for i, s := range shares {
			bytesShares[i], _ = json.Marshal(s)
		}

		// 序列化
		jsonBytes, err := json.Marshal(bytesShares)
		if err != nil {
			fmt.Println("Error marshaling []Share:", err)
			return nil, nil
		}
		fmt.Println("==================secret", jsonBytes)

		fmt.Printf("Serialized: %s\n", jsonBytes)
		// 将字节数组转换为 Base64 编码的字符串
		bytesBase64 := base64.StdEncoding.EncodeToString(jsonBytes)
		fmt.Println("===============", bytesBase64)
		return &types.SecretSSRes{
			Shares: string(bytesBase64),
		}, nil
	case "2":
		k := req.K
		shares1 := req.Shares
		fmt.Println(shares1)
		// string -> []byte
		shares, err := base64.StdEncoding.DecodeString(shares1)
		fmt.Println("==================shares", shares)
		// 反序列化
		var deserializedBytes [][]byte
		err = json.Unmarshal(shares, &deserializedBytes)
		if err != nil {
			fmt.Println("Error unmarshaling []Share:", err)
		}

		// 将 []byte 转换回 []Share
		deserializedShares := make([]ss.Share, len(deserializedBytes))
		for i, b := range deserializedBytes {
			err := json.Unmarshal(b, &deserializedShares[i])
			if err != nil {
				fmt.Println("Error unmarshaling Share:", err)
			}
		}

		// 调用 ss.Combine 函数，传入切片中的所有 shares
		secret1 := ss.Combine(k, deserializedShares...)
		// 将字节数组转换为 Base64 编码的字符串
		bytesBase64 := base64.StdEncoding.EncodeToString(secret1)
		fmt.Println("===============", bytesBase64)
		return &types.SecretSSRes{
			Secret: bytesBase64,
		}, nil
	}
	return
}
