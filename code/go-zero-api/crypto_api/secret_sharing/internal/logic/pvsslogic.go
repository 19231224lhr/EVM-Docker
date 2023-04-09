package logic

import (
	"blockchain-crypto/crypto_api/secret_sharing/internal/svc"
	"blockchain-crypto/crypto_api/secret_sharing/internal/types"
	"blockchain-crypto/secret_sharing/pvss"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Secured-Finance/kyber"
	"github.com/Secured-Finance/kyber/proof/dleq"
	"github.com/Secured-Finance/kyber/share"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type PVSSLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPVSSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PVSSLogic {
	return &PVSSLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// x
func serializeScalars(scalars []kyber.Scalar) (string, error) {
	bytesArray := make([][]byte, len(scalars))
	for i, s := range scalars {
		bytes, err := s.MarshalBinary()
		if err != nil {
			return "", err
		}
		bytesArray[i] = bytes
	}

	jsonBytes, err := json.Marshal(bytesArray)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func deserializeScalars(jsonStr string, suite kyber.Group) ([]kyber.Scalar, error) {
	var bytesArray [][]byte
	err := json.Unmarshal([]byte(jsonStr), &bytesArray)
	if err != nil {
		return nil, err
	}

	scalars := make([]kyber.Scalar, len(bytesArray))
	for i, bytes := range bytesArray {
		s := suite.Scalar()
		err := s.UnmarshalBinary(bytes)
		if err != nil {
			return nil, err
		}
		scalars[i] = s
	}

	return scalars, nil
}

// X
func serializePoints(points []kyber.Point) (string, error) {
	bytesArray := make([][]byte, len(points))
	for i, p := range points {
		bytes, err := p.MarshalBinary()
		if err != nil {
			return "", err
		}
		bytesArray[i] = bytes
	}

	jsonBytes, err := json.Marshal(bytesArray)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func deserializePoints(jsonStr string, suite kyber.Group) ([]kyber.Point, error) {
	var bytesArray [][]byte
	err := json.Unmarshal([]byte(jsonStr), &bytesArray)
	if err != nil {
		return nil, err
	}

	points := make([]kyber.Point, len(bytesArray))
	for i, bytes := range bytesArray {
		p := suite.Point()
		err := p.UnmarshalBinary(bytes)
		if err != nil {
			return nil, err
		}
		points[i] = p
	}

	return points, nil
}

// secret
func serializeScalar(scalar kyber.Scalar) (string, error) {
	bytes, err := scalar.MarshalBinary()
	if err != nil {
		return "", err
	}

	jsonBytes, err := json.Marshal(bytes)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func deserializeScalar(jsonStr string, suite kyber.Group) (kyber.Scalar, error) {
	var bytes []byte
	err := json.Unmarshal([]byte(jsonStr), &bytes)
	if err != nil {
		return nil, err
	}

	scalar := suite.Scalar()
	err = scalar.UnmarshalBinary(bytes)
	if err != nil {
		return nil, err
	}

	return scalar, nil
}

// encShares
type SerializedPubVerShare struct {
	SI int
	SV []byte
	PC []byte
	PR []byte
	VG []byte
	VH []byte
}

func serializePubVerShares(shares []*pvss.PubVerShare) (string, error) {
	var serializedShares []SerializedPubVerShare

	for _, share := range shares {
		svBytes, err := share.S.V.MarshalBinary()
		if err != nil {
			return "", err
		}

		pcBytes, err := share.P.C.MarshalBinary()
		if err != nil {
			return "", err
		}

		prBytes, err := share.P.R.MarshalBinary()
		if err != nil {
			return "", err
		}

		vgBytes, err := share.P.VG.MarshalBinary()
		if err != nil {
			return "", err
		}

		vhBytes, err := share.P.VH.MarshalBinary()
		if err != nil {
			return "", err
		}

		serializedShares = append(serializedShares, SerializedPubVerShare{
			SI: share.S.I,
			SV: svBytes,
			PC: pcBytes,
			PR: prBytes,
			VG: vgBytes,
			VH: vhBytes,
		})
	}

	jsonBytes, err := json.Marshal(serializedShares)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func deserializePubVerShares(jsonStr string, suite kyber.Group) ([]*pvss.PubVerShare, error) {
	var serializedShares []SerializedPubVerShare
	err := json.Unmarshal([]byte(jsonStr), &serializedShares)
	if err != nil {
		return nil, err
	}

	var shares []*pvss.PubVerShare
	for _, s := range serializedShares {
		sv := suite.Point()
		err = sv.UnmarshalBinary(s.SV)
		if err != nil {
			return nil, err
		}

		pc := suite.Scalar()
		err = pc.UnmarshalBinary(s.PC)
		if err != nil {
			return nil, err
		}

		pr := suite.Scalar()
		err = pr.UnmarshalBinary(s.PR)
		if err != nil {
			return nil, err
		}

		vg := suite.Point()
		err = vg.UnmarshalBinary(s.VG)
		if err != nil {
			return nil, err
		}

		vh := suite.Point()
		err = vh.UnmarshalBinary(s.VH)
		if err != nil {
			return nil, err
		}

		shares = append(shares, &pvss.PubVerShare{
			S: share.PubShare{
				I: s.SI,
				V: sv,
			},
			P: dleq.Proof{
				C:  pc,
				R:  pr,
				VG: vg,
				VH: vh,
			},
		})
	}

	return shares, nil
}

// pubPloy

type ExportedPubPoly struct {
	G       string `json:"g"`
	B       string `json:"b"`
	Commits string `json:"commits"`
}

func pubPolyToExported(suite kyber.Group, pubPoly *share.PubPoly) *ExportedPubPoly {
	g := suite.String() // 假设 suite 具有 String() 方法
	b, commits := pubPoly.Info()

	// 将 b 和 commits 转换为字符串
	bStr, err := serializePoint(b)
	B, err := deserializePoint(bStr, suite)
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%% bstr : ", bStr)
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%% : ", B.String() == b.String())
	if err != nil {
		fmt.Println("err in pubPolyToExported serializePoint", err)
	}
	commitsStr, err := serializePoints(commits)
	if err != nil {
		fmt.Println("Error in pubPolyToExported !!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return nil
	}

	return &ExportedPubPoly{
		G:       g,
		B:       bStr,
		Commits: commitsStr,
	}
}

func exportedToPubPoly(suite kyber.Group, exported *ExportedPubPoly) (*share.PubPoly, error) {
	// 将字符串转换回 g、b 和 commits
	fmt.Println("&&&&&&&&&&&&&&& exported.B = ", exported.B)
	B, err := deserializePoint(exported.B, suite)
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& B : ", B)
	if err != nil {
		fmt.Println("error in exportedToPubPoly B", err)
	}

	g := suite

	commits, err := deserializePoints(exported.Commits, suite)
	if err != nil {
		fmt.Println("err in exportedToPubPoly 2 : ", err)
		return nil, err
	}

	return share.NewPubPoly(g, B, commits), nil
}

// H
func serializePoint(point kyber.Point) (string, error) {
	bytes, err := point.MarshalBinary()
	if err != nil {
		return "", err
	}
	jsonBytes, err := json.Marshal(bytes)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func deserializePoint(jsonStr string, suite kyber.Group) (kyber.Point, error) {
	var bytes []byte
	err := json.Unmarshal([]byte(jsonStr), &bytes)
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!error in deserializePoint 1 : ", err)
	}
	point := suite.Point()
	err = point.UnmarshalBinary(bytes)
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!!!!error in deserializePoint 2 : ", err)
	}
	return point, nil
}

func (l *PVSSLogic) PVSS(req *types.SecretPvssReq) (resp *types.SecretPvssRes, err error) {
	// todo: add your logic here and delete this line

	fmt.Println("=============START SECRET_SHARING PVSS=============")
	//设置公开参数，包括生成元G,H以及选取的椭圆曲线等信息suite
	suite, G, H := pvss.Init() // H maybe different
	fmt.Printf("======== suite = %v, G = %v, H = %v\n", suite, G, H)

	switch req.Name {
	case "1":
		fmt.Println("=====================Case 1=======================")
		n := req.N
		var x []kyber.Scalar
		var X []kyber.Point
		var secret kyber.Scalar
		x, X = pvss.CreateKeyPair(suite, n)
		secret = pvss.CreateSecret(suite)
		fmt.Printf("************* x = %v, X = %v, secret = %v\n", x, X, secret)

		// 将 x 转换为 JSON
		// jsonBytes_x, err := json.Marshal(x)
		jsonBytes_x, err := serializeScalars(x)
		if err != nil {
			fmt.Println("Error marshaling x:", err)
		}
		jsonx := string(jsonBytes_x)
		fmt.Println("======================x to JSON====================:", jsonx)

		// 将 X 转换为 JSON
		// 将 X 转换为 JSON
		jsonX, err := serializePoints(X)
		if err != nil {
			fmt.Println("Error serializing X:", err)
		}
		fmt.Println("======================X JSON====================:", jsonX)

		// 将 secret 转换为 JSON
		// 将 secret 转换为 JSON
		jsonsecret, err := serializeScalar(secret)
		if err != nil {
			fmt.Println("Error serializing Scalar:", err)
		}
		fmt.Println("======================secret JSON====================:", jsonsecret)

		// H -> json
		jsonH, err := serializePoint(H)
		if err != nil {
			fmt.Printf("Error serializing pointH: %v\n", err)
		}

		fmt.Printf("Serialized pointH: %s\n", jsonH)
		return &types.SecretPvssRes{
			PubX:   jsonX,
			Prix:   jsonx,
			Secret: jsonsecret,
			H:      jsonH,
		}, nil
	case "2":
		fmt.Println("=====================Case 2=======================")
		X := req.X
		secret := req.Secret
		t := req.T
		// X : json -> X
		// 将 JSON 字符串转换回 X
		deserializedX, err := deserializePoints(X, suite)
		if err != nil {
			fmt.Println("Error deserializing x:", err)
		}
		// secret : json -> secret
		// 将 JSON 字符串转换回 secret
		deserializedsecret, err := deserializeScalar(secret, suite)
		if err != nil {
			fmt.Println("Error deserializing secret:", err)
		}

		// json -> H
		fmt.Println("***********************************************")
		fmt.Println("+++++++++++++++++++++++++++++++++", req.H)
		H1, err := deserializePoint(req.H, suite)
		fmt.Printf("******************Deserialized H: %v\n", H1)

		//////////////////////
		encShares, pubPoly, _ := pvss.CreateEnShare(suite, H1, deserializedX, deserializedsecret, t)
		///////////////////////////

		fmt.Println("enShares = ", encShares)
		fmt.Println("pubPoly = ", pubPoly)

		// enShares -> json
		jsonStr_enShares, err := serializePubVerShares(encShares)
		if err != nil {
			fmt.Println("Error serializing PubVerShares:", err)
		}

		fmt.Println("============Serialized PubVerShares============:", jsonStr_enShares)
		// pubPoly -> json
		exported_pubpoly := pubPolyToExported(suite, pubPoly)
		fmt.Println("pubpoly to json struct : ", exported_pubpoly)
		fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBB : ", exported_pubpoly.B)
		// 序列化为 JSON
		jsonData_pubpoly, err := json.Marshal(exported_pubpoly)
		if err != nil {
			log.Fatal("Error marshaling ExportedPubPoly:", err)
		}
		var b = string(jsonData_pubpoly)
		fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC : ", b)

		return &types.SecretPvssRes{
			EncShares: jsonStr_enShares,
			PubPoly:   string(jsonData_pubpoly),
		}, nil
	case "3":
		fmt.Println("=====================Case 3=======================")
		X := req.X
		encShares_string := req.EncShares
		pubPoly_string := req.PubPoly

		// json -> enshares
		// 使用 'deserializePubVerShares' 函数将 JSON 字符串反序列化为 pvss.PubVerShare 列表
		deserializedShares, err := deserializePubVerShares(encShares_string, suite)
		if err != nil {
			fmt.Println("Error deserializing PubVerShares:", err)
		}
		fmt.Printf("==============True enShares: %v\n", deserializedShares)
		///////////////////////////////////////
		// json -> pubPoly
		// 反序列化 JSON
		var unmarshalled ExportedPubPoly
		err = json.Unmarshal([]byte(pubPoly_string), &unmarshalled)
		if err != nil {
			log.Fatal("Error unmarshaling ExportedPubPoly:", err)
		}
		// 转换回原始结构体
		restoredPubPoly, err := exportedToPubPoly(suite, &unmarshalled)
		if err != nil {
			log.Fatal("Error converting ExportedPubPoly back to PubPoly:", err)
		}
		fmt.Println("========True pubpoly: ", restoredPubPoly)
		///////////////////////////////////////////
		// X : json -> X
		// 将 JSON 字符串转换回 X
		deserializedX, err := deserializePoints(X, suite)
		if err != nil {
			fmt.Println("Error deserializing x:", err)
		}

		// H : json -> H
		H1, err := deserializePoint(req.H, suite)
		if err != nil {
			fmt.Printf("Error deserializing H1: %v\n", err)
		}
		fmt.Printf("****************Deserialized point: %v\n", H1)

		//  (2) 验证（任何人）:给定分享者公钥X[1]，多项式的承诺信息pubPoly，加密的分享信息encShares[1]，生成元H，验证加密的秘密分享信息的一致性
		res := pvss.VerifyEncShare(suite, H1, deserializedX[1], restoredPubPoly, deserializedShares[1])
		fmt.Println(res)

		// json -> x
		// 将 JSON 字符串转换回 []kyber.Scalar 类型
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx : ", req.Xpri)
		x, err := deserializeScalars(req.Xpri, suite)
		if err != nil {
			fmt.Println("Error deserializing x:", err)
		}
		// 打印反序列化后的 x
		fmt.Println("========True x: ", x)

		return &types.SecretPvssRes{
			Res: res,
		}, nil
	case "4":
		fmt.Println("=====================Case 4=======================")
		encShares_json := req.EncShares
		pubPoly_json := req.PubPoly
		X_json := req.X
		x_json := req.Xpri
		H_json := req.H
		N := req.N
		T := req.T

		// json -> enshares
		// 使用 'deserializePubVerShares' 函数将 JSON 字符串反序列化为 pvss.PubVerShare 列表
		enShares, err := deserializePubVerShares(encShares_json, suite)
		if err != nil {
			fmt.Println("Error deserializing PubVerShares:", err)
		}
		fmt.Printf("==============True enShares: %v\n", enShares)

		// json -> pubPoly
		// 反序列化 JSON
		var unmarshalled ExportedPubPoly
		err = json.Unmarshal([]byte(pubPoly_json), &unmarshalled)
		if err != nil {
			log.Fatal("Error unmarshaling ExportedPubPoly:", err)
		}
		// 转换回原始结构体
		pubPoly, err := exportedToPubPoly(suite, &unmarshalled)
		if err != nil {
			log.Fatal("Error converting ExportedPubPoly back to PubPoly:", err)
		}
		fmt.Println("========True pubpoly: ", pubPoly)

		// X : json -> X
		// 将 JSON 字符串转换回 X
		X, err := deserializePoints(X_json, suite)
		if err != nil {
			fmt.Println("Error deserializing x:", err)
		}
		fmt.Println("========True X: ", X)

		// json -> x
		// 将 JSON 字符串转换回 []kyber.Scalar 类型
		fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx : ", x_json)
		x, err := deserializeScalars(x_json, suite)
		if err != nil {
			fmt.Println("Error deserializing x:", err)
		}
		// 打印反序列化后的 x
		fmt.Println("========True x: ", x)

		// H : json -> H
		H1, err := deserializePoint(H_json, suite)
		if err != nil {
			fmt.Printf("Error deserializing H1: %v\n", err)
		}
		fmt.Printf("========True H: %v\n", H1)

		fmt.Printf("\nlen(X) = %v, len(enShares) = %v\n", len(X), len(enShares))

		// start
		// （4）检查解密的分享并尽可能恢复机密（分享者/第三方）（默认会进行一致性检验）
		//参与者公钥列表
		var K []pvss.KyberPoint
		//参与者加密分享信息列表
		var E []*pvss.PubVerShare
		//参与者解密后的子秘密列表
		var D []*pvss.PubVerShare

		for i := 0; i < N; i++ {
			K = append(K, X[i])
			E = append(E, enShares[i])
		}
		var a *pvss.PubVerShare
		for i := 1; i <= 3; i++ {
			a, _ = pvss.DecShare(suite, H, X[i], x[i], enShares[i], pubPoly)
			D = append(D, a)
			fmt.Println("==============a : ", a)
		}
		fmt.Println("==============D : ", D)
		recovered, _ := pvss.RecoverSecret(suite, G, K, E, D, T, N)
		fmt.Println("==============recovered : ", recovered)

		// H -> json
		Recovered, err := serializePoint(recovered)
		if err != nil {
			fmt.Printf("Error serializing Recovered: %v\n", err)
		}

		return &types.SecretPvssRes{
			Recovered: Recovered,
		}, nil
	case "5":
		//设置公开参数，包括生成元G,H以及选取的椭圆曲线等信息suite
		suite, G, H := pvss.Init()

		//设置参与人数n和门限值t
		n := 5
		t := 2

		//创建秘密分享者的公私钥对
		x, X := pvss.CreateKeyPair(suite, n)

		// 创建秘密分享值secret
		secret := pvss.CreateSecret(suite)

		// （1） 秘密分配（分发者）：给定曲线信息suite，生成元H，获得加密的秘密分享信息列表encShares和对多项式的承诺信息pubPoly
		encShares, pubPoly, _ := pvss.CreateEnShare(suite, H, X, secret, t)

		//  (2) 验证（任何人）:给定分享者公钥X[1]，多项式的承诺信息pubPoly，加密的分享信息encShares[1]，生成元H，验证加密的秘密分享信息的一致性
		res := pvss.VerifyEncShare(suite, H, X[1], pubPoly, encShares[1])
		fmt.Println(res)
		fmt.Println("suite = ", suite)

		// （3） 解密获得未加密的子秘密（分享者）（默认会进行一致性检验）
		decShare1, _ := pvss.DecShare(suite, H, X[1], x[1], encShares[1], pubPoly)
		fmt.Println(decShare1)
		decShare2, _ := pvss.DecShare(suite, H, X[2], x[2], encShares[2], pubPoly)
		decShare3, _ := pvss.DecShare(suite, H, X[3], x[3], encShares[3], pubPoly)

		// （4）检查解密的分享并尽可能恢复机密（分享者/第三方）（默认会进行一致性检验）
		//参与者公钥列表
		var K []pvss.KyberPoint
		//参与者加密分享信息列表
		var E []*pvss.PubVerShare
		//参与者解密后的子秘密列表
		var D []*pvss.PubVerShare

		for i := 1; i < 4; i++ {
			K = append(K, X[i])
			E = append(E, encShares[i])
		}
		D = append(D, decShare1)
		D = append(D, decShare2)
		D = append(D, decShare3)

		// 秘密恢复结果recovered
		recovered, _ := pvss.RecoverSecret(suite, G, K, E, D, t, n)
		fmt.Println(recovered)
	default:
		fmt.Println("wrong number")
	}
	return
}
