package main

import (
	"blockchain-crypto/key_exchange/dh"
	"fmt"
)

func main() {

	// 使用2048位群
	group := dh.Init_4096()

	// Alice创建公私钥对
	alicePrivate, alicePublic, err := group.GenerateKey()
	if err != nil {
		fmt.Printf("Failed to generate alice's private / public key pair: %s", err)
	}

	// Bob创建公私钥对
	bobPrivate, bobPublic, err := group.GenerateKey()
	if err != nil {
		fmt.Printf("Failed to generate bob's private / public key pair: %s", err)
	}

	//Alice计算会话密钥
	secretAlice := group.ComputeSecret(alicePrivate, bobPublic)

	//Bob计算会话密钥
	secretBob := group.ComputeSecret(bobPrivate, alicePublic)
	fmt.Println("12312312312")
	fmt.Println(secretAlice == secretBob)
	fmt.Println(secretBob)

}
