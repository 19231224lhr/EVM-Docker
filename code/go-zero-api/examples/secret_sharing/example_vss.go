package main

import (
	"blockchain-crypto/secret_sharing/vss"
	"crypto/elliptic"
	"fmt"
	"math/big"
)

func main() {

	//选定曲线ec，设置参与者num，门限值t
	ec := elliptic.P384()
	num := 6
	t := 3

	//设置秘密值secret，
	secret := new(big.Int).SetInt64(1000)

	//给定曲线ec,门限值t,秘密值secret，参与者num获得对秘密分享多项式的承诺vs和秘密分享结果share，以及错误信息err
	vs, shares, _ := vss.Create(ec, t, secret, num)
	// fmt.Println(shares[0].ID)

	//给定曲线ec，门限值t，秘密分享多项式的承诺vs，对子秘密share[0]验证
	res := shares[0].Verify(ec, t, vs)
	fmt.Println(res)

	//给定秘密分享者的子集subset(使之包含share[1],share[2],share[3])，曲线ec，重新恢复原秘密secret1
	var subset vss.Shares
	subset = append(subset, shares[1])
	subset = append(subset, shares[2])
	subset = append(subset, shares[3])
	secret1, _ := subset.ReConstruct(ec)
	fmt.Println(secret1)

}
