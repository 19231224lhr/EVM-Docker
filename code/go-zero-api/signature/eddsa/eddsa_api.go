package eddsa

import "blockchain-crypto/signature/eddsa/dependency"

// 所有api函数的输入输出均为由字节数组转化成的字符串（验证输出是bool）

// 密钥生成
// 随机生成公私钥对并输出
// 输出格式：私钥，公钥
func Keygen_api() (string, string) {
	pk, sk, _ := dependency.GenerateKey(nil)
	return string(sk), string(pk)
}

// 签名生成
// 对指定消息进行签名，输入消息将使用sha512进行hash，无需预先hash
// 输入格式：私钥，消息
// 输出格式：签名
func Sign_api(seck string, mes string) string {
	sk := []byte(seck)
	m := []byte(mes)
	sig := dependency.Sign(sk, m)
	return string(sig)
}

// 签名验证
// 对消息签名进行验证
// 输入格式：公钥，消息，签名
// 输出格式：是否正确（0表示否，1表示是）
func Verify_api(pubk string, mes string, sig string) bool {
	pk := []byte(pubk)
	m := []byte(mes)
	sign := []byte(sig)
	return dependency.Verify(pk, m, sign)
}
