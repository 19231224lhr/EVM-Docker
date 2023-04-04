package vss

import (
	"crypto/elliptic"
	"errors"
	"fmt"
	"github.com/bnb-chain/tss-lib/common"
	"github.com/bnb-chain/tss-lib/crypto"
	"math/big"
)

type (
	Share struct {
		Threshold int //门限值
		ID,       //x点坐标
		Share *big.Int //y点坐标
	}

	Vs []*crypto.ECPoint // 多项式的承诺值，对应于一系列椭圆曲线上的点

	Shares []*Share //秘密分享切片
)

var (
	ErrNumSharesBelowThreshold = fmt.Errorf("not enough shares to satisfy the threshold")

	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

// 检查Shamir的秘密共享的共享ID，如果发现重复或0值，则返回错误
func CheckIndexes(ec elliptic.Curve, indexes []*big.Int) ([]*big.Int, error) {
	visited := make(map[string]struct{})
	for _, v := range indexes {
		vMod := new(big.Int).Mod(v, ec.Params().N)
		if vMod.Cmp(zero) == 0 {
			return nil, errors.New("party index should not be 0")
		}
		vModStr := vMod.String()
		if _, ok := visited[vModStr]; ok {
			return nil, fmt.Errorf("duplicate indexes %s", vModStr)
		}
		visited[vModStr] = struct{}{}
	}
	return indexes, nil
}

// 返回由Shamir的秘密共享算法创建的新的秘密共享数组，需要从输入的秘密中重新创建最小数量的共享（长度共享）
func Create(ec elliptic.Curve, threshold int, secret *big.Int, num int) (Vs, Shares, error) {

	threshold = threshold - 1
	ids := make([]*big.Int, 0)
	for i := 0; i < num; i++ {
		ids = append(ids, common.GetRandomPositiveInt(ec.Params().N))
	}

	if threshold < 1 {
		return nil, nil, errors.New("vss threshold < 1")
	}

	ids, err := CheckIndexes(ec, ids)
	if err != nil {
		return nil, nil, err
	}

	if num < threshold {
		return nil, nil, ErrNumSharesBelowThreshold
	}

	poly := samplePolynomial(ec, threshold, secret)
	poly[0] = secret
	v := make(Vs, len(poly))
	for i, ai := range poly {
		v[i] = crypto.ScalarBaseMult(ec, ai)
	}

	shares := make(Shares, num)
	for i := 0; i < num; i++ {
		share := evaluatePolynomial(ec, threshold, poly, ids[i])
		shares[i] = &Share{Threshold: threshold, ID: ids[i], Share: share}
	}
	return v, shares, nil
}

// 验证子秘密的有效性
func (share *Share) Verify(ec elliptic.Curve, threshold int, vs Vs) bool {
	threshold = threshold - 1
	if share.Threshold != threshold || vs == nil {
		return false
	}
	var err error
	modQ := common.ModInt(ec.Params().N)
	v, t := vs[0], one
	for j := 1; j <= threshold; j++ {

		t = modQ.Mul(t, share.ID)

		vjt := vs[j].SetCurve(ec).ScalarMult(t)
		v, err = v.SetCurve(ec).Add(vjt)
		if err != nil {
			return false
		}
	}
	sigmaGi := crypto.ScalarBaseMult(ec, share.Share)
	return sigmaGi.Equals(v)
}

// 恢复原秘密
func (shares Shares) ReConstruct(ec elliptic.Curve) (secret *big.Int, err error) {
	if shares != nil && shares[0].Threshold > len(shares) {
		return nil, ErrNumSharesBelowThreshold
	}
	modN := common.ModInt(ec.Params().N)

	xs := make([]*big.Int, 0)
	for _, share := range shares {
		xs = append(xs, share.ID)
	}

	secret = zero
	for i, share := range shares {
		times := one
		for j := 0; j < len(xs); j++ {
			if j == i {
				continue
			}
			sub := modN.Sub(xs[j], share.ID)
			subInv := modN.ModInverse(sub)
			div := modN.Mul(xs[j], subInv)
			times = modN.Mul(times, div)
		}

		fTimes := modN.Mul(share.Share, times)
		secret = modN.Add(secret, fTimes)
	}

	return secret, nil
}

func samplePolynomial(ec elliptic.Curve, threshold int, secret *big.Int) []*big.Int {
	q := ec.Params().N
	v := make([]*big.Int, threshold+1)
	v[0] = secret
	for i := 1; i <= threshold; i++ {
		ai := common.GetRandomPositiveInt(q)
		v[i] = ai
	}
	return v
}

func evaluatePolynomial(ec elliptic.Curve, threshold int, v []*big.Int, id *big.Int) (result *big.Int) {
	q := ec.Params().N
	modQ := common.ModInt(q)
	result = new(big.Int).Set(v[0])
	X := big.NewInt(int64(1))
	for i := 1; i <= threshold; i++ {
		ai := v[i]
		X = modQ.Mul(X, id)
		aiXi := new(big.Int).Mul(ai, X)
		result = modQ.Add(result, aiXi)
	}
	return
}
