package main

import (
	"fmt"
	"math/big"

	"blockchain-crypto/commitment/polynomial_commitment"
)

func main() {

	//创建公开参数T
	_, T := polynomial_commitment.Init("hello", 10)

	//定义多项式的系数
	var p []*big.Int
	p = append(p, new(big.Int).SetInt64(20))
	p = append(p, new(big.Int).SetInt64(2))
	p = append(p, new(big.Int).SetInt64(20))
	p = append(p, new(big.Int).SetInt64(2))

	//创建承诺
	r := polynomial_commitment.Commit(p, T)
	
	//验证承诺
	u := polynomial_commitment.Verify(r, p, T)
	fmt.Println(u)

	//创建p（x）在x=1处的见证
	d := new(big.Int).SetInt64(1)
	X, Y, W := polynomial_commitment.CreatWitness(T, p, d)

	// 验证见证
	v := polynomial_commitment.VerifyWitness(T, r, X, Y, W)
	fmt.Println(v)
}
