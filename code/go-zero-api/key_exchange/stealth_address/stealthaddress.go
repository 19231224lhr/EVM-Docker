package stealth_address

import (
	"crypto/elliptic"

	"crypto/sha256"
	"math/big"
)

//公钥的结构，X/Y分别为椭圆曲线上的点的两个坐标。
type PublicKey struct {
	X *big.Int
	Y *big.Int
}

//接收方计算公私钥对，str1,str2字符串相当于输入的随机因子。
func RecCalculateKeyPairs(str1, str2 string) (priv1 *big.Int, pub1 *PublicKey, priv2 *big.Int, pub2 *PublicKey) {

	com1 := sha256.Sum256([]byte(str1))
	k1 := com1[:]
	priv1 = new(big.Int).SetBytes(k1)

	com2 := sha256.Sum256([]byte(str1))
	k2 := com2[:]
	priv2 = new(big.Int).SetBytes(k2)

	pub1 = new(PublicKey)
	pub2 = new(PublicKey)

	pub1.X, pub1.Y = elliptic.P256().ScalarBaseMult(k1)
	pub2.X, pub2.Y = elliptic.P256().ScalarBaseMult(k2)

	return priv1, pub1, priv2, pub2
}

//发送方利用接收方的公钥以及自身的私钥计算隐地址P。
func SendCalculateObfuscateAddress(r *big.Int, pub1, pub2 *PublicKey) (P *PublicKey) {

	rA := new(PublicKey)
	rA.X, rA.Y = elliptic.P256().ScalarMult(pub1.X, pub1.Y, r.Bytes())

	rAxbyte := rA.X.Bytes()
	rAybyte := rA.Y.Bytes()

	for i := 0; i < len(rAybyte); i++ {
		rAxbyte = append(rAxbyte, rAybyte[i])
	}
	HashrA := sha256.Sum256(rAxbyte)
	t := HashrA[:]
	HashrAG := new(PublicKey)

	HashrAG.X, HashrAG.Y = elliptic.P256().ScalarBaseMult(t)

	P = new(PublicKey)
	P.X, P.Y = elliptic.P256().Add(HashrAG.X, HashrAG.Y, pub2.X, pub2.Y)

	return P
}

//发送方计算自身的私钥r对应的公钥R。
func SendCalculatePublicKey(r *big.Int) (R *PublicKey) {

	R = new(PublicKey)
	rb := r.Bytes()
	R.X, R.Y = elliptic.P256().ScalarBaseMult(rb)

	return R
}

//接收方通过隐地址P和公钥R，利用自身私钥计算这笔交易是否转向自身。
func RecCalculateObfuscateAddress(P, R *PublicKey, priv1 *big.Int, pub2 *PublicKey) bool {

	aR := new(PublicKey)
	aR.X, aR.Y = elliptic.P256().ScalarMult(R.X, R.Y, priv1.Bytes())

	aRxbyte := aR.X.Bytes()
	aRybyte := aR.Y.Bytes()

	for i := 0; i < len(aRybyte); i++ {
		aRxbyte = append(aRxbyte, aRybyte[i])
	}
	HashaR := sha256.Sum256(aRxbyte)

	t := HashaR[:]
	HashaRG := new(PublicKey)

	HashaRG.X, HashaRG.Y = elliptic.P256().ScalarBaseMult(t)

	P1 := new(PublicKey)
	P1.X, P1.Y = elliptic.P256().Add(HashaRG.X, HashaRG.Y, pub2.X, pub2.Y)

	return P1.X.String() == P.X.String() && P1.Y.String() == P.Y.String()
}

//接收方确认P,R对应的交易接收方为自身时，利用该函数计算P对应的私钥以此来花费这笔钱。
func RecCalculateAddressPrivatekey(priv1, priv2 *big.Int, R *PublicKey) (priv *big.Int) {

	aR := new(PublicKey)
	aR.X, aR.Y = elliptic.P256().ScalarMult(R.X, R.Y, priv1.Bytes())

	aRxbyte := aR.X.Bytes()
	aRybyte := aR.Y.Bytes()

	for i := 0; i < len(aRybyte); i++ {
		aRxbyte = append(aRxbyte, aRybyte[i])
	}
	HashaR := sha256.Sum256(aRxbyte)

	priv = new(big.Int)
	priv.SetBytes(HashaR[:])
	priv.Add(priv, priv2)

	return priv
}
