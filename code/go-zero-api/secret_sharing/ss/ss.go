package ss

import (
	"github.com/Psiphon-Labs/psiphon-tunnel-core/psiphon/common/sss"
)

type Share struct {
	X byte
	Y []byte
}

func Split(n byte, k byte, secret []byte) (shares []Share) {
	var share Share
	res, _ := sss.Split(n, k, secret)
	for x, y := range res {
		share.X = x
		share.Y = y
		shares = append(shares, share)
	}
	return shares
}

func Combine(k byte, shares ...Share) []byte {

	subset := make(map[byte][]byte, k)

	for i := 0; i < int(k); i++ {
		x := shares[i].X
		y := shares[i].Y
		subset[x] = y
	}

	return sss.Combine(subset)
}
