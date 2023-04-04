package main

import (
	"blockchain-crypto/secret_sharing/ss"
	"encoding/base64"
	"fmt"
)

func main() {

	//设定参与者人数n，门限值k
	n := byte(10)
	k := byte(3)

	//设定秘密值secret
	secret := []byte{1, 2, 3, 4, 4, 5, 4, 5, 100, 200}

	// 将字节数组转换为 Base64 编码的字符串
	bytesBase64 := base64.StdEncoding.EncodeToString(secret)
	fmt.Println("===============", bytesBase64)

	string_byte, err := base64.StdEncoding.DecodeString(bytesBase64)
	if err != nil {
		return
	}
	fmt.Println("===============", string_byte)

	a := string(secret)
	fmt.Println("===============", []byte(a))
	//给定参与者人数n，门限值k，秘密值secret，获得秘密分享结果shares
	shares := ss.Split(n, k, secret)

	//给定门限值k，秘密分享结果的子集{shares[1],shares[2],shares[3]}恢复原秘密secret1
	secret1 := ss.Combine(k, shares[1], shares[2], shares[3])
	fmt.Println(secret1)
}
