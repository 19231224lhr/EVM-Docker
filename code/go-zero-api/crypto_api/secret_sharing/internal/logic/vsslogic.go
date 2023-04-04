package logic

import (
	"blockchain-crypto/crypto_api/secret_sharing/internal/svc"
	"blockchain-crypto/crypto_api/secret_sharing/internal/types"
	"blockchain-crypto/secret_sharing/vss"
	"context"
	"crypto/elliptic"
	"encoding/json"
	"fmt"
	"github.com/bnb-chain/tss-lib/tss"
	"github.com/zeromicro/go-zero/core/logx"
	"math/big"
)

type VSSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVSSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VSSLogic {
	return &VSSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
type Vs []*ECPoint

type ECPoint struct {
	Curve  elliptic.Curve
	Coords [2]*big.Int
}

type ecPointJSON struct {
	Curve string `json:"curve"`
	X     string `json:"x"`
	Y     string `json:"y"`
}

func (vs Vs) MarshalJSON() ([]byte, error) {
	points := make([]ecPointJSON, len(vs))
	for i, point := range vs {
		points[i] = ecPointJSON{
			Curve: point.Curve.Params().Name,
			X:     point.Coords[0].String(),
			Y:     point.Coords[1].String(),
		}
	}
	return json.Marshal(points)
}

func (vs *Vs) UnmarshalJSON(data []byte) error {
	points := make([]ecPointJSON, 0)
	err := json.Unmarshal(data, &points)
	if err != nil {
		return err
	}

	*vs = make(Vs, len(points))
	for i, point := range points {
		x := new(big.Int)
		y := new(big.Int)
		x.SetString(point.X, 10)
		y.SetString(point.Y, 10)

		var curve elliptic.Curve
		switch point.Curve {
		case "P-256":
			curve = elliptic.P256()
		case "P-384":
			curve = elliptic.P384()
		case "P-521":
			curve = elliptic.P521()
		default:
			return fmt.Errorf("unsupported curve: %s", point.Curve)
		}

		(*vs)[i] = &ECPoint{
			Curve: curve,
			Coords: [2]*big.Int{
				x,
				y,
			},
		}
	}
	return nil
}

*/

// ===================================
// type Shares []*Share
//
//	type Share struct {
//		Threshold int
//		ID, Share *big.Int
//	}
//
//	type shareJSON struct {
//		Threshold int     `json:"threshold"`
//		ID        big.Int `json:"id"`
//		Share     big.Int `json:"share"`
//	}
//
//	func (s Shares) MarshalJSON() ([]byte, error) {
//		shares := make([]shareJSON, len(s))
//		for i, share := range s {
//			shares[i] = shareJSON{
//				Threshold: share.Threshold,
//				ID:        *share.ID,
//				Share:     *share.Share,
//			}
//		}
//		return json.Marshal(shares)
//	}
//
//	func (s *Shares) UnmarshalJSON(data []byte) error {
//		fmt.Printf("Received JSON data: %s\n", data)
//		var sharesJSON []shareJSON
//		err := json.Unmarshal(data, &sharesJSON)
//		fmt.Println("&&&&&&&&&&&&&&&&&************************1")
//		if err != nil {
//			fmt.Printf("Error during Unmarshal: %v\n", err)
//			return err
//		}
//		fmt.Println("&&&&&&&&&&&&&&&&&************************5")
//		shares := make([]*Share, len(sharesJSON))
//		fmt.Println("&&&&&&&&&&&&&&&&&************************2")
//		for i, sj := range sharesJSON {
//			fmt.Println("&&&&&&&&&&&&&&&&&************************3")
//			fmt.Println("&&&&&&&&&&&&&&&&&************************4", reflect.TypeOf(sj.ID))
//			id := sj.ID
//
//			share := sj.Share
//
//			shares[i] = &Share{
//				Threshold: sj.Threshold,
//				ID:        &id,
//				Share:     &share,
//			}
//		}
//
//		*s = shares
//		return nil
//	}
//
// 在程序启动时向 tss 包注册名为 "P-384" 的 NIST P-384 椭圆曲线。这样，在之后的代码中，你可以使用这个曲线名称来执行与椭圆曲线密码相关的操作
func init() {
	curve := elliptic.P384()
	tss.RegisterCurve("P-384", curve)
}

func (l *VSSLogic) VSS(req *types.SecretVssReq) (resp *types.SecretVssRes, err error) {
	// todo: add your logic here and delete this line
	switch req.Name {
	case "1":
		fmt.Println("========================Case 1==========================")
		secretnum := req.SecretNum
		var secret *big.Int
		secret = new(big.Int).SetInt64(secretnum)
		// 将 *big.Int 转换为字符串
		bigIntStr := secret.String()
		fmt.Println("BigInt as string:", bigIntStr)

		return &types.SecretVssRes{
			Secret: bigIntStr,
		}, nil

	case "2":
		fmt.Println("========================Case 2==========================")
		ec := elliptic.P384()
		num := req.Num
		t := req.T
		var secret *big.Int
		// 将字符串转换回 *big.Int
		newBigInt := new(big.Int)
		_, success := newBigInt.SetString(req.Secret, 10)
		if !success {
			fmt.Println("Error converting string back to *big.Int")
			return
		}
		fmt.Println("BigInt restored from string:", newBigInt)
		secret = newBigInt
		var vs vss.Vs
		var shares vss.Shares
		vs, shares, _ = vss.Create(ec, t, secret, num) // 分享多项式的承诺vs和秘密分享结果share
		fmt.Println("======================Vs True====================:", vs)
		fmt.Println("====================Shares True==================:", shares)
		// 将 Vs 转换为 JSON
		jsonBytes, err := json.Marshal(vs)
		if err != nil {
			fmt.Println("Error marshaling Vs:", err)
		}
		jsonVs := string(jsonBytes)
		fmt.Println("======================Vs JSON====================:", jsonVs)
		// shares -> JSON
		jsonBytes1, err := json.Marshal(shares)
		if err != nil {
			panic(err)
		}
		jsonShares := string(jsonBytes1)
		fmt.Println("====================shares JSON==================:", jsonShares)
		// Test shares[0].Verify() function
		return &types.SecretVssRes{
			Vs:     jsonVs,
			Shares: jsonShares,
		}, nil
	case "3":
		fmt.Println("========================Case 3==========================")
		//给定曲线ec，门限值t，秘密分享多项式的承诺vs，对子秘密share[0]验证
		ec := elliptic.P384()
		t := req.T
		var vs_string string
		vs_string = req.Vs
		var shares_string string
		shares_string = req.Shares
		fmt.Printf("=======================Value======================== : ec = %v, t = %v, vs = %v, shares = %v\n", ec, t, vs_string, shares_string)
		// 将 VS_JSON 转换回 Vs
		fmt.Println("=======================Vs String======================:", vs_string)
		var vs vss.Vs
		err = json.Unmarshal([]byte(vs_string), &vs)
		if err != nil {
			fmt.Println("Error unmarshaling Vs:", err)
		}
		fmt.Println("========================Vs True=======================:", vs)
		// 将 Shares_JSON 转换回 shares
		fmt.Println("=====================Shares String====================:", shares_string)
		var shares vss.Shares
		err = json.Unmarshal([]byte(shares_string), &shares)
		if err != nil {
			fmt.Println("Error unmarshaling Shares:", err)
		}
		fmt.Println("=====================Shares True====================:", shares)
		// shares -> JSON
		jsonBytes1, err := json.Marshal(shares)
		if err != nil {
			panic(err)
		}
		jsonShares := string(jsonBytes1)
		fmt.Println("====================shares JSON==================:", jsonShares)
		fmt.Println("====================Is Json Same?==================:", jsonShares == shares_string)
		if jsonShares != shares_string {
			panic("error: json is not same")
		}
		//给定曲线ec，门限值t，秘密分享多项式的承诺vs，对子秘密share[0]验证
		res := shares[0].Verify(ec, t, vs)
		fmt.Println("=====================Finally====================:", res)
		return &types.SecretVssRes{
			Res: res,
		}, nil
	case "4":
		fmt.Println("========================Case 4==========================")
		//给定秘密分享者的子集subset(使之包含share[1],share[2],share[3])，曲线ec，重新恢复原秘密secret1
		ec := elliptic.P384()
		t := req.T
		shares_json := req.Shares
		var shares vss.Shares
		err = json.Unmarshal([]byte(shares_json), &shares)
		if err != nil {
			fmt.Println("Error unmarshaling Shares:", err)
		}
		fmt.Println("=====================Shares True====================:", shares)
		var subset vss.Shares
		for i := 1; i <= t; i++ {
			subset = append(subset, shares[i])
		}
		secret, _ := subset.ReConstruct(ec)
		fmt.Println("========================Secret======================:", secret)
		// secret -> json
		jsonBytes, err := json.Marshal(secret)
		if err != nil {
			panic(err)
		}
		jsonSecret := string(jsonBytes)
		fmt.Println("====================Secret JSON==================:", jsonSecret)
		return &types.SecretVssRes{
			Secret1: jsonSecret,
		}, nil
	default:
		fmt.Println("num is wrong, not in 1-4, your input num is ", req.Num)
	}
	return
}
