package logic

import (
	"blockchain-crypto/crypto_api/vrf/internal/svc"
	"blockchain-crypto/crypto_api/vrf/internal/types"
	"blockchain-crypto/vrf"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type VrfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVrfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VrfLogic {
	return &VrfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VrfLogic) Vrf(req *types.VrfReq) (resp *types.VrfRes, err error) {
	// todo: add your logic here and delete this line
	//生成公私钥对(priv,pub)
	priv, _ := vrf.Newprivatekey()
	pub, _ := vrf.GeneratePublickey(priv)
	fmt.Println("priv : ", priv)
	fmt.Println("pub : ", pub)
	// may make a map to copy pub
	switch req.Name {
	case "1":
		//利用消息message和私钥priv获得vrf结果vrf0以及证明proof
		message := []byte(req.Message)
		vrf0, proof, _ := vrf.Prove(priv, message)

		// 将vrf0字节数组转换为 Base64 编码的字符串
		vrf0byte := base64.StdEncoding.EncodeToString(vrf0)
		fmt.Println("===============", vrf0byte)
		// 将proof字节数组转换为 Base64 编码的字符串
		proofbyte := base64.StdEncoding.EncodeToString(proof)
		fmt.Println("===============", proofbyte)
		return &types.VrfRes{
			Vrf0:  vrf0byte,
			Proof: proofbyte,
		}, nil
	case "2":
		vrf0, err := base64.StdEncoding.DecodeString(req.Vrf0)
		if err != nil {
			fmt.Println("error : ", err)
		}
		fmt.Println("===============", vrf0)
		proof, err := base64.StdEncoding.DecodeString(req.Proof)
		if err != nil {
			fmt.Println("error : ", err)
		}
		fmt.Println("===============", proof)
		//利用公钥pub，消息message，证明proof 验证vrf结果vrf0的正确性
		res, _ := vrf.Verify(pub, []byte(req.Message), vrf0, proof)
		fmt.Println(res)
		return &types.VrfRes{
			Res: res,
		}, nil
	default:
		fmt.Println("wrong num")
		return nil, nil
	}
}
