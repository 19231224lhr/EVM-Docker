package ecdh
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//显示ECC初始参数，包括所选取的椭圆曲线类型，生成元点G，阶数N，模P
func Init() {
	fmt.Printf("--ECC Parameters--\n")
	fmt.Printf(" Name: %s\n", elliptic.P256().Params().Name)
	fmt.Printf(" N: %x\n", elliptic.P256().Params().N)
	fmt.Printf(" P: %x\n", elliptic.P256().Params().P)
	fmt.Printf(" Gx: %x\n", elliptic.P256().Params().Gx)
	fmt.Printf(" Gy: %x\n", elliptic.P256().Params().Gy)
}

//计算公私钥对（priv, pub）(在椭圆曲线p256上)
func CalculateKeypair() (priv *big.Int, pub ecdsa.PublicKey) {
	a, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	priv = a.D
	pub = a.PublicKey
	return priv, pub
}

//计算协商后的密钥(shared)
func CalculateNegotiationKey(pub ecdsa.PublicKey, priv *big.Int) (shared [32]byte) {
	a, _ := pub.Curve.ScalarMult(pub.X, pub.Y, priv.Bytes())
	shared = sha256.Sum256(a.Bytes())
	return shared
}
