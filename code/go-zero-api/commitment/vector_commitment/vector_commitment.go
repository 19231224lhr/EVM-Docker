package vector_commitment

import (
	"crypto/sha256"
	"math/big"

	"github.com/cloudflare/bn256"
)

// 公开参数的结构体
type Trust struct {
	T1i  []*bn256.G1
	T1ij [][]*bn256.G1
	T2i  []*bn256.G2
	T2ij [][]*bn256.G2
}

// 利用str字符串作为随机因子生成公开参数，n为多项式的最大次数，返回秘密值c和公开参数T。
func Init(str string, n int) (c []byte, T *Trust) {
	T = new(Trust)
	var r []*big.Int
	var rij [][]*big.Int
	com := sha256.Sum256([]byte(str))
	for i := 0; i < n; i++ {
		r = append(r, new(big.Int).SetBytes(com[:]))
		com = sha256.Sum256([]byte(com[:]))
	}

	for i := 0; i < n; i++ {
		var tmp []*big.Int
		for j := 0; j < n; j++ {
			t := new(big.Int).Mul(r[i], r[j])
			tmp = append(tmp, new(big.Int).Mod(t, bn256.Order))
		}
		rij = append(rij, tmp)
	}

	for i := 0; i < n; i++ {

		t1 := new(bn256.G1).ScalarBaseMult(r[i])
		T.T1i = append(T.T1i, t1)
	}

	for i := 0; i < n; i++ {
		t2 := new(bn256.G2).ScalarBaseMult(r[i])
		T.T2i = append(T.T2i, t2)
	}

	for i := 0; i < n; i++ {
		var tmp []*bn256.G1
		for j := 0; j < n; j++ {
			t1 := new(bn256.G1).ScalarBaseMult(rij[i][j])
			tmp = append(tmp, t1)
		}
		T.T1ij = append(T.T1ij, tmp)
	}

	for i := 0; i < n; i++ {
		var tmp []*bn256.G2
		for j := 0; j < n; j++ {
			t2 := new(bn256.G2).ScalarBaseMult(rij[i][j])
			tmp = append(tmp, t2)
		}
		T.T2ij = append(T.T2ij, tmp)
	}
	return c, T
}

//创建对向量p的承诺值
func Commit(p []*big.Int, T *Trust) (r *bn256.G1) {
	var u []*bn256.G1

	for i := 0; i < len(p); i++ {
		t := new(bn256.G1).ScalarMult(T.T1i[i], p[i])
		u = append(u, t)
	}

	r = u[0]

	for i := 1; i < len(p); i++ {
		r = new(bn256.G1).Add(r, u[i])
	}

	return r
}

//创建向量p在i处值为I的证明。
func CreatWitness(i int, p []*big.Int, T *Trust) (I *big.Int, W *bn256.G1) {
	I = p[i]
	var u []*bn256.G1
	for j := 0; j < len(p); j++ {
		if i != j {
			u = append(u, new(bn256.G1).ScalarMult(T.T1ij[i][j], p[j]))
		}
	}

	W = u[0]

	for i := 1; i < len(u); i++ {
		W = new(bn256.G1).Add(W, u[i])
	}
	return I, W
}

//验证证明。
func VerifyWitness(i int, I *big.Int, T *Trust, W *bn256.G1, r *bn256.G1) bool {

	h := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	himi := new(bn256.G1).ScalarMult(T.T1i[i], I)
	himiNEG := new(bn256.G1).Neg(himi)
	Chimi := new(bn256.G1).Add(r, himiNEG)
	e1 := bn256.Pair(Chimi, T.T2i[i])
	e2 := bn256.Pair(W, h)

	return e1.String() == e2.String()
}
