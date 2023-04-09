package main

import (
	"fmt"

	"blockchain-crypto/secret_sharing/pvss"
)

func main() {

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
}
