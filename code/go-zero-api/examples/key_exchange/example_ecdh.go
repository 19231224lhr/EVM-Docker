package main

import (
	"blockchain-crypto/key_exchange/ecdh"
	"fmt"
)

func main() {
	//需要密钥协商的双方各自计算公私钥对（privatekey1, publickey1）,(privatekey2, publickey2),其中公钥是椭圆曲线上由G和私钥生成的点的坐标
	privatekey1, publickey1 := ecdh.CalculateKeypair()
	privatekey2, publickey2 := ecdh.CalculateKeypair()

	//需要密钥协商的双方获得对方的公钥后并利用自身的私钥计算协商后的密钥
	shared1 := ecdh.CalculateNegotiationKey(publickey2, privatekey1)
	shared2 := ecdh.CalculateNegotiationKey(publickey1, privatekey2)

	fmt.Printf("%x\n", shared1)
	fmt.Printf("%x\n", shared2)
}
