package main

import (
	"encoding/json"
	"fmt"
	"github.com/Secured-Finance/kyber"
	"github.com/Secured-Finance/kyber/proof/dleq"
	"github.com/Secured-Finance/kyber/share"
	"log"

	"blockchain-crypto/secret_sharing/pvss"
)

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

type ExportedPubPoly struct {
	G       string `json:"g"`
	B       string `json:"b"`
	Commits string `json:"commits"`
}

func pubPolyToExported(suite kyber.Group, pubPoly *share.PubPoly) *ExportedPubPoly {
	g := suite.String() // 假设 suite 具有 String() 方法
	b, commits := pubPoly.Info()

	// 将 b 和 commits 转换为字符串
	bStr := b.String()
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
	pointb := suite.Point()
	err := pointb.UnmarshalBinary([]byte(exported.B))
	if err != nil {
		fmt.Println("err in exportedToPubPoly 1")
	}
	g := suite

	commits, err := deserializePoints(exported.Commits, suite)
	if err != nil {
		fmt.Println("err in exportedToPubPoly 2")
		return nil, err
	}

	return share.NewPubPoly(g, pointb, commits), nil
}

func main() {
	fmt.Println("test")
	//设置公开参数，包括生成元G,H以及选取的椭圆曲线等信息suite
	suite, G, H := pvss.Init()

	//设置参与人数n和门限值t
	n := 5
	t := 2

	//创建秘密分享者的公私钥对
	x, X := pvss.CreateKeyPair(suite, n)

	// 创建秘密分享值secret
	secret := pvss.CreateSecret(suite)

	// print x
	fmt.Println("======================x True=================", x)

	// 将 x 转换为 JSON
	// jsonBytes_x, err := json.Marshal(x)
	jsonBytes_x, err := serializeScalars(x)
	if err != nil {
		fmt.Println("Error marshaling x:", err)
	}
	jsonx := string(jsonBytes_x)
	fmt.Println("======================x to JSON====================:", jsonx)

	// 将 JSON 字符串转换回 []kyber.Scalar 类型
	deserializedx, err := deserializeScalars(jsonBytes_x, suite)
	if err != nil {
		fmt.Println("Error deserializing x:", err)
	}

	// 打印反序列化后的 x
	fmt.Println("======================Json to x====================:", deserializedx)

	fmt.Println("======================X True=================", X)
	// 将 X 转换为 JSON
	jsonX, err := serializePoints(X)
	if err != nil {
		fmt.Println("Error serializing X:", err)
	}
	fmt.Println("======================X JSON====================:", jsonX)

	// 将 JSON 字符串转换回 X
	deserializedX, err := deserializePoints(jsonX, suite)
	if err != nil {
		fmt.Println("Error deserializing x:", err)
	}

	// 打印反序列化后的 X
	fmt.Println("======================json to X====================:", deserializedX)

	fmt.Println("======================secret True=================", secret)
	// 将 secret 转换为 JSON
	jsonsecret, err := serializeScalar(secret)
	if err != nil {
		fmt.Println("Error serializing Scalar:", err)
	}
	fmt.Println("======================secret JSON====================:", jsonsecret)

	// 将 JSON 字符串转换回 secret
	deserializedsecret, err := deserializeScalar(jsonsecret, suite)
	if err != nil {
		fmt.Println("Error deserializing secret:", err)
	}
	// 打印反序列化后的 secret
	fmt.Println("======================json to secret====================:", deserializedsecret)

	// （1） 秘密分配（分发者）：给定曲线信息suite，生成元H，获得加密的秘密分享信息列表encShares和对多项式的承诺信息pubPoly
	encShares, pubPoly, _ := pvss.CreateEnShare(suite, H, X, secret, t)

	// g kyber.Group is suite
	// pubPoly.Info() return b kyber.Point, commits []kyber.Point

	// encShares
	fmt.Println("======================encShares True=================", encShares)
	// 使用 'serializePubVerShares' 函数将 pvss.PubVerShare 列表序列化为 JSON 字符串
	jsonStr, err := serializePubVerShares(encShares)
	if err != nil {
		fmt.Println("Error serializing PubVerShares:", err)
		return
	}

	fmt.Println("Serialized PubVerShares:", jsonStr)

	// 使用 'deserializePubVerShares' 函数将 JSON 字符串反序列化为 pvss.PubVerShare 列表
	deserializedShares, err := deserializePubVerShares(jsonStr, suite)
	if err != nil {
		fmt.Println("Error deserializing PubVerShares:", err)
		return
	}

	fmt.Printf("Deserialized PubVerShares: %v\n", deserializedShares)

	// start test pubPoly ===================================================
	fmt.Println("====================pubPoly===================", pubPoly)
	// 转换为可导出的结构体
	exported_pubpoly := pubPolyToExported(suite, pubPoly)
	fmt.Println("pubpoly to json struct : ", exported_pubpoly)
	// 序列化为 JSON
	jsonData_pubpoly, err := json.Marshal(exported_pubpoly)
	if err != nil {
		log.Fatal("Error marshaling ExportedPubPoly:", err)
	}
	fmt.Println("========pubpoly to json : ", string(jsonData_pubpoly))
	// 反序列化 JSON
	var unmarshalled ExportedPubPoly
	err = json.Unmarshal(jsonData_pubpoly, &unmarshalled)
	if err != nil {
		log.Fatal("Error unmarshaling ExportedPubPoly:", err)
	}
	fmt.Println("=======json to json struct", jsonData_pubpoly)
	// 转换回原始结构体
	restoredPubPoly, err := exportedToPubPoly(suite, &unmarshalled)
	fmt.Println("========json  pubpoly: ", restoredPubPoly)
	if err != nil {
		log.Fatal("Error converting ExportedPubPoly back to PubPoly:", err)
	}

	//  (2) 验证（任何人）:给定分享者公钥X[1]，多项式的承诺信息pubPoly，加密的分享信息encShares[1]，生成元H，验证加密的秘密分享信息的一致性
	res := pvss.VerifyEncShare(suite, H, X[1], pubPoly, encShares[1])
	fmt.Println(res)

	// （3） 解密获得未加密的子秘密（分享者）（默认会进行一致性检验）
	decShare1, _ := pvss.DecShare(suite, H, X[1], x[1], encShares[1], pubPoly)
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
	fmt.Println(secret)
	fmt.Println(recovered)
}
