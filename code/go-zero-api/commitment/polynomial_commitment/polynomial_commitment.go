package polynomial_commitment

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"github.com/cloudflare/bn256"

	"math/big"
)

//公开参数的结构体
type Trust struct {
	T1 []*bn256.G1
	T2 []*bn256.G2
}

//利用str字符串作为随机因子生成公开参数，n为多项式的最大次数，返回秘密值c和公开参数T。
func Init(str string, n int) (c []byte, T *Trust) {
	var r []*big.Int
	com := sha256.Sum256([]byte(str))
	c = com[:]
	k := new(big.Int).SetBytes(c)
	t := new(big.Int).SetInt64(1)

	T = new(Trust)

	for i := 0; i < n; i++ {
		r = append(r, new(big.Int).Mod(t, bn256.Order))
		t.Mul(t, k)
	}

	for i := 0; i < n; i++ {

		t1 := new(bn256.G1).ScalarBaseMult(r[i])
		T.T1 = append(T.T1, t1)
	}

	for i := 0; i < n; i++ {
		t2 := new(bn256.G2).ScalarBaseMult(r[i])
		T.T2 = append(T.T2, t2)
	}

	return c, T
}

//对多项式p进行承诺。
func Commit(p []*big.Int, T *Trust) (r *bn256.G1) {
	var u []*bn256.G1

	for i := 0; i < len(p); i++ {
		t := new(bn256.G1).ScalarMult(T.T1[i], p[i])
		u = append(u, t)
	}

	r = u[0]

	for i := 1; i < len(p); i++ {
		r = new(bn256.G1).Add(r, u[i])
	}

	return r
}

//验证承诺
func Verify(r *bn256.G1, p []*big.Int, T *Trust) bool {
	r1 := Commit(p, T)

	return r1.String() == r.String()

}

func div(p []*big.Int, d *big.Int) (q []*big.Int) {
	for i := 0; i < len(p)-1; i++ {
		q = append(q, new(big.Int).SetInt64(0))
	}
	res := new(big.Int)

	for i := len(p) - 1; i > 0; i-- {
		deepCopyByGob(p[i], q[i-1])
		res.Mul(d, p[i])
		p[i-1].Add(p[i-1], res)
		// fmt.Println(q[i-1])
	}

	return q
}
func deepCopyByGob(src, dst interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buffer).Decode(dst)
}

//为多项式在y=p(x)处创建证明W。
func CreatWitness(T *Trust, p []*big.Int, d *big.Int) (X, Y *big.Int, W *bn256.G1) {

	X = d
	var u []*big.Int
	var p1 []*big.Int

	for i := 0; i < len(p); i++ {
		u = append(u, new(big.Int).SetInt64(0))
	}
	for i := 0; i < len(p); i++ {
		p1 = append(p1, new(big.Int).SetInt64(0))
	}

	for i := 0; i < len(p); i++ {
		deepCopyByGob(p[i], p1[i])
	}

	Y = p1[0]
	for i := 1; i < len(p); i++ {

		u[i] = p1[i]
		for j := 0; j < i; j++ {
			u[i].Mul(u[i], X)
		}
		Y = Y.Add(Y, u[i])

	}
	// fmt.Println(Y)

	q := div(p, d)
	W = Commit(q, T)
	return X, Y, W
}

//验证证明是否正确。
func VerifyWitness(T *Trust, r *bn256.G1, X, Y *big.Int, W *bn256.G1) bool {

	s2 := T.T2[1]

	zG2Neg := new(bn256.G2).Neg(new(bn256.G2).ScalarBaseMult(X))

	sz := new(bn256.G2).Add(s2, zG2Neg)

	yG1Neg := new(bn256.G1).Neg(new(bn256.G1).ScalarBaseMult(Y))

	cy := new(bn256.G1).Add(r, yG1Neg)
	h := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	e1 := bn256.Pair(W, sz)
	e2 := bn256.Pair(cy, h)
	return e1.String() == e2.String()
}
