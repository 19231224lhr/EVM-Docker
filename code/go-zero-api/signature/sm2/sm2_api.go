package SM2

import (
	"blockchain-crypto/signature/sm2/dependency"
	"crypto/rand"
	"math/big"
)

// 密钥生成
// 随机生成公私钥对并输出
// 输出格式：私钥，公钥
// 私钥是10进制整数，公钥是字节数组
func Keygen_api() (string, string) {
	sk, _ := dependency.GenerateKey(nil)
	pk := dependency.Compress(&sk.PublicKey)
	return sk.D.String(), string(pk)
}

// 签名生成
// 对指定消息进行签名，输入消息将使用sha512进行hash，无需预先hash
// 输入格式：私钥，消息
// 输出格式：签名
func Sign_api(seck string, mes string) string {
	m := []byte(mes)
	var temp big.Int
	temp.SetString(seck, 10)
	c := dependency.P256Sm2()
	pkx, pky := c.ScalarBaseMult(temp.Bytes())
	sk := &dependency.PrivateKey{
		PublicKey: dependency.PublicKey{
			Curve: c,
			X:     pkx,
			Y:     pky,
		},
		D: &temp,
	}
	sig, _ := sk.Sign(rand.Reader, m, nil)
	return string(sig)
}

// 签名验证
// 对消息签名进行验证
// 输入格式：公钥，消息，签名
// 输出格式：是否正确（0表示否，1表示是）
func Verify_api(pubk string, mes string, sig string) bool {
	pk := dependency.Decompress([]byte(pubk))
	m := []byte(mes)
	sign := []byte(sig)
	return pk.Verify(m, sign)
}
