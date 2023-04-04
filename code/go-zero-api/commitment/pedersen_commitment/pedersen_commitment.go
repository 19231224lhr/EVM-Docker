package pedersen_commitment

import (
	"github.com/bwesterb/go-ristretto"
)

// ParamsGen 生成两个曲线上的点，用于计算承诺
func ParamsGen() (G, H ristretto.Point) {
	G.Rand()
	H.Rand()
	return G, H
}

// RandomGen 生成一个随机阶数
func RandomGen() (r ristretto.Scalar) {
	r.Rand()
	return r
}

// Commit 计算秘密secret的Pedersen承诺，返回曲线上的点
func Commit(G, H ristretto.Point, s uint64, r ristretto.Scalar) (commit ristretto.Point) {
	var x ristretto.Scalar
	x.SetUint64(s)
	//c = xG + rH
	var comm ristretto.Point
	comm.Add(G.ScalarMult(&G, &x), H.ScalarMult(&H, &r))
	return comm
}

// Verify 通过原始参数验证承诺
func Verify(commit, G, H ristretto.Point, s uint64, r ristretto.Scalar) bool {
	var x ristretto.Scalar
	x.SetUint64(s)
	var calculateComm ristretto.Point
	calculateComm.Add(G.ScalarMult(&G, &x), H.ScalarMult(&H, &r))
	return calculateComm.Equals(&commit)
}
