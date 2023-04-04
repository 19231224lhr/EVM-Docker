package ECDSA

import (
	"blockchain-crypto/signature/ecdsa/dependency"
)

// 所有api函数的输入输出均为由字节数组转化成的字符串（验证输出是bool）
// 使用序列化（Serialize）函数将结构体转化为字节数组，使用解析（Parse）函数将字节数组转化为结构体

// 密钥生成
// 随机生成公私钥对并输出
// 输出格式：私钥，公钥
func Keygen_api() (string, string) {
	sk, _ := dependency.NewPrivateKey(dependency.S256())
	pk := sk.PubKey()
	return string(sk.Serialize()), string(pk.SerializeCompressed())
}

// 签名生成
// 对指定消息进行签名，输入消息应先进行hash运算再输入
// 输入格式：私钥，消息
// 输出格式：签名
func Sign_api(seck string, meshashed string) string {
	sk, _ := dependency.PrivKeyFromBytes(dependency.S256(), []byte(seck))
	sig, _ := sk.Sign([]byte(meshashed))
	return string(sig.Serialize())
}

// 签名验证
// 对消息签名进行验证
// 输入格式：公钥，消息，签名
// 输出格式：是否正确（0表示否，1表示是）
func Verify_api(pubk string, meshashed string, sig string) bool {
	pk, _ := dependency.ParsePubKey([]byte(pubk), dependency.S256())
	sign, _ := dependency.ParseSignature([]byte(sig), dependency.S256())
	return sign.Verify([]byte(meshashed), pk)
}
