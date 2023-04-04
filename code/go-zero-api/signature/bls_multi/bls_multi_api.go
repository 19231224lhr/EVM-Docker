package BLS_Multi

import (
	"blockchain-crypto/signature/bls/dependency/bls"
	"blockchain-crypto/signature/bls_multi/dependency/groupsig"
)

// Chinese notes encode by UTF-8

// 所有api函数的输入输出均为16进制（验证输出是bool）
// 需要输入输出多个时使用数组

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
// 对指定消息进行签名
// 输入格式：私钥，消息
// 输出格式：签名
func Sign_api(seck string, mes string) string {
	bls.Init(bls.CurveFp254BNb)
	var sk bls.SecretKey
	sk.SetHexString(seck)
	sign := sk.Sign(mes)
	return sign.GetHexString()
}

// 签名整合
// 将多个签名整合为一个多重签名
// 输入格式：签名数组
// 输出格式：签名
func MultiSign_api(sigs []string) string {
	bls.Init(bls.CurveFp254BNb)
	var signs []groupsig.Signature
	signs = make([]groupsig.Signature, len(sigs))
	for i := 0; i < len(sigs); i++ {
		signs[i].SetHexString(sigs[i])
	}
	mulsig := groupsig.AggregateSigs(signs)
	return mulsig.GetHexString()
}

// 多重签名验证
// 将多个公钥整合为一个公钥，再验证多重签名
// 输入格式：公钥数组，消息，多重签名
// 输出格式：签名是否有效
func MultiVerify_api(pubks []string, mes string, mulsig string) bool {
	bls.Init(bls.CurveFp254BNb)
	var pks []groupsig.Pubkey
	pks = make([]groupsig.Pubkey, len(pubks))
	for i := 0; i < len(pubks); i++ {
		pks[i].SetHexString(pubks[i])
	}
	m := []byte(mes)
	var msig groupsig.Signature
	msig.SetHexString(mulsig)
	return groupsig.VerifyAggregateSig(pks, m, msig)
}
