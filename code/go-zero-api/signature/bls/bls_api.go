package bls

import "C"
import (
	"blockchain-crypto/signature/bls/dependency/bls"
)

// 所有api函数的输入输出均为16进制（验证输出是bool）

// 密钥生成
// 随机生成公私钥对并输出
// 输出格式：私钥，公钥
func Keygen_api() (string, string) {
	bls.Init(bls.CurveFp254BNb)
	var sk bls.SecretKey
	sk.SetByCSPRNG()
	pk := sk.GetPublicKey()
	return sk.GetHexString(), pk.GetHexString()
}

// 签名生成
// 对指定消息进行签名，输入消息使用内置的hash函数进行hash运算
// 输入格式：私钥，消息
// 输出格式：签名
func Sign_api(seck string, mes string) string {
	bls.Init(bls.CurveFp254BNb)
	var sk bls.SecretKey
	sk.SetHexString(seck)
	sign := sk.Sign(mes)
	return sign.GetHexString()
}

// 签名验证
// 对消息签名进行验证
// 输入格式：公钥，消息，签名
// 输出格式：是否正确（0表示否，1表示是）
func Verify_api(pubk string, mes string, sig string) bool {
	bls.Init(bls.CurveFp254BNb)
	var pk bls.PublicKey
	pk.SetHexString(pubk)
	var sign bls.Sign
	sign.SetHexString(sig)
	return sign.Verify(&pk, mes)
}
