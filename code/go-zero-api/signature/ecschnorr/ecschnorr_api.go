package ECSchnorr

import (
	"blockchain-crypto/signature/ecschnorr/dependency"
	"crypto/sha256"
	"math/big"
)

// 密钥生成
// 随机生成公私钥对并输出
// 输出格式：私钥，公钥
// 均为字节数组转化为字符串
func Keygen_api() (string, string) {
	schnorr := dependency.NewSchnorr(dependency.S256(), sha256.New())
	seck, _ := schnorr.KeyGen()
	pub := seck.PublicKey
	return string(dependency.FromECDSA(seck)), string(dependency.FromECDSAPub(&pub))
}

// 签名生成
// 对指定消息进行签名，输入消息无需预先hash
// 输入格式：私钥，消息
// 输出格式：签名r，签名s
// 输入转为字节数组处理
// 输出为10进制字符串
func Sign_api(seck string, mes string) (string, string) {
	schnorr := dependency.NewSchnorr(dependency.S256(), sha256.New())
	sk, _ := dependency.ToECDSA([]byte(seck))
	sigr, sigs, _ := schnorr.Sign(sk, []byte(mes))
	return sigr.String(), sigs.String()
}

// 签名验证
// 对消息签名进行验证
// 输入格式：公钥，消息，签名r，签名s
// 输出格式：是否正确（0表示否，1表示是）
// 输入的公钥和消息转为字节数组，签名r和s转为10进制整数处理
func Verify_api(pubk string, mes string, sigr string, sigs string) bool {
	schnorr := dependency.NewSchnorr(dependency.S256(), sha256.New())
	pk, _ := dependency.UnmarshalPubkey([]byte(pubk))
	var r, s big.Int
	r.SetString(sigr, 10)
	s.SetString(sigs, 10)
	result, _ := schnorr.Verify(pk, &r, &s, []byte(mes))
	return result
}
