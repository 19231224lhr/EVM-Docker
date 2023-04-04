package bls_threshold

import "blockchain-crypto/signature/bls/dependency/bls"

// 私钥秘密分享
// 对私钥进行秘密分享
// 输入格式：私钥，门限值，总份额数
// 输出格式：ID数组，私钥数组，公钥数组
func SecKeyShare_api(seck string, t int, n int) ([]string, []string, []string) {
	bls.Init(bls.CurveFp254BNb)
	var sk bls.SecretKey
	sk.SetHexString(seck)
	msk := sk.GetMasterSecretKey(t)
	mpk := bls.GetMasterPublicKey(msk)
	idVec := make([]bls.ID, n)
	secVec := make([]bls.SecretKey, n)
	pubVec := make([]bls.PublicKey, n)
	idVecStr := make([]string, n)
	secVecStr := make([]string, n)
	pubVecStr := make([]string, n)
	for i := 0; i < n; i++ {
		idVec[i].SetLittleEndian([]byte{byte(i & 255), byte(i >> 8), 2, 3, 4, 5})
		idVecStr[i] = idVec[i].GetHexString()
		secVec[i].Set(msk, &idVec[i])
		secVecStr[i] = secVec[i].GetHexString()
		pubVec[i].Set(mpk, &idVec[i])
		pubVecStr[i] = pubVec[i].GetHexString()
	}
	return idVecStr, secVecStr, pubVecStr
}

// 私钥恢复
// 从多个私钥份额中恢复出私钥
// 输入格式：ID数组，私钥数组
// 输出格式：是否正确运行，恢复的私钥
func SecKeyRecover_api(idVector []string, secVector []string) (bool, string) {
	bls.Init(bls.CurveFp254BNb)
	if len(idVector) != len(secVector) {
		return false, "incompatible length"
	}
	k := len(idVector)
	idVec := make([]bls.ID, k)
	secVec := make([]bls.SecretKey, k)
	for i := 0; i != k; i++ {
		idVec[i].SetHexString(idVector[i])
		secVec[i].SetHexString(secVector[i])
	}
	var sec bls.SecretKey
	_ = sec.Recover(secVec, idVec)
	secKeyReStr := sec.GetHexString()
	return true, secKeyReStr
}

// 签名恢复
// 从多个签名份额中恢复出签名
// 输入格式：ID数组，签名数组
// 输出格式：是否正确运行，恢复的签名
func SigRecover_api(idVector []string, sigVector []string) (bool, string) {
	bls.Init(bls.CurveFp254BNb)
	if len(idVector) != len(sigVector) {
		return false, "incompatible length"
	}
	k := len(idVector)
	idVec := make([]bls.ID, k)
	sigVec := make([]bls.Sign, k)
	for i := 0; i != k; i++ {
		idVec[i].SetHexString(idVector[i])
		sigVec[i].SetHexString(sigVector[i])
	}
	var sig bls.Sign
	_ = sig.Recover(sigVec, idVec)
	sigReStr := sig.GetHexString()
	return true, sigReStr
}

// 公钥恢复
// 从多个公钥份额中恢复出公钥
// 输入格式：ID数组，公钥数组
// 输出格式：是否正确运行，恢复的公钥
func PubKeyRecover_api(idVector []string, pubVector []string) (bool, string) {
	bls.Init(bls.CurveFp254BNb)
	if len(idVector) != len(pubVector) {
		return false, "incompatible length"
	}
	k := len(idVector)
	idVec := make([]bls.ID, k)
	pubVec := make([]bls.PublicKey, k)
	for i := 0; i != k; i++ {
		idVec[i].SetHexString(idVector[i])
		pubVec[i].SetHexString(pubVector[i])
	}
	var pub bls.PublicKey
	_ = pub.Recover(pubVec, idVec)
	pubKeyReStr := pub.GetHexString()
	return true, pubKeyReStr
}
