package bls_purego

import (
	bls "blockchain-crypto/signature/bls_purego/dependency"
	"crypto/rand"
)

// 所有api函数的输入输出均为字节串转换为字符串（验证输出是bool）

// 密钥生成
// 随机生成公私钥对并输出
// 输出格式：私钥，公钥
func Keygen_api() (string, string) {
	var sk bls.PrivateKey
	seed := make([]byte, 32)
	rand.Read(seed)
	sk = bls.KeyGen(seed)
	pk := sk.GetPublicKey()
	return string(sk.Bytes()), string(pk.Bytes())
}

// 签名生成
// 对指定消息进行签名，输入消息使用内置的hash函数进行hash运算
// 输入格式：私钥，消息
// 输出格式：签名
func Sign_api(seck string, mes string) string {
	asm := new(bls.AugSchemeMPL)
	var sk bls.PrivateKey
	sk = bls.KeyFromBytes([]byte(seck))
	sign := asm.Sign(sk, []byte(mes))
	return string(sign)
}

// 签名验证
// 对消息签名进行验证
// 输入格式：公钥，消息，签名
// 输出格式：是否正确（0表示否，1表示是）
func Verify_api(pubk string, mes string, sig string) bool {
	asm := new(bls.AugSchemeMPL)
	var pk bls.PublicKey
	pk, _ = bls.NewPublicKey([]byte(pubk))
	var sign []byte
	sign = []byte(sig)
	return asm.Verify(pk, []byte(mes), sign)
}
