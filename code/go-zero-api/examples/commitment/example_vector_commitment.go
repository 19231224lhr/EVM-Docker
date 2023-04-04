package main

import (
	"fmt"
	"math/big"

	"blockchain-crypto/commitment/vector_commitment"
)

func main() {

	//创建公开参数T（最大允许对10维向量进行承诺）
	_, T := vector_commitment.Init("hello", 10)

	//定义向量
	var p []*big.Int
	p = append(p, new(big.Int).SetInt64(3))
	p = append(p, new(big.Int).SetInt64(55))
	p = append(p, new(big.Int).SetInt64(44))
	p = append(p, new(big.Int).SetInt64(45))
	p = append(p, new(big.Int).SetInt64(99))
	p = append(p, new(big.Int).SetInt64(20))

	//创建承诺
	r := vector_commitment.Commit(p, T)

	//创建向量p[3]处值为I的证明。
	I, W := vector_commitment.CreatWitness(3, p, T)

	//验证证明
	v := vector_commitment.VerifyWitness(3, I, T, W, r)
	fmt.Println(v)
}
