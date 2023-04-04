package dependency

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

var (
	// kh     = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	// pri, _ = crypto.HexToECDSA(kh)

	msg = Keccak256([]byte("xueyang.han@okcoin.com"))
)

func Test_Schnorr(t *testing.T) {
	schnorr := NewSchnorr(S256(), sha256.New())
	// schnorr := NewSchnorr(nil, nil)
	pri, err := schnorr.KeyGen()
	if err != nil {
		t.Error(err)
	}

	t.Logf("key:%s", hex.EncodeToString(FromECDSA(pri)))

	msg := []byte("@xueyang.han")
	r, s, err := schnorr.Sign(pri, msg)
	if err != nil {
		t.Error(err)
	}
	// t.Log(hex.EncodeToString(sig), len(sig)) //view

	pub := &pri.PublicKey
	valid, err := schnorr.Verify(pub, r, s, msg)
	if err != nil {
		t.Error(err)
	}
	if !valid {
		t.Error(ErrSigCheck)
	}
}

// var (
// 	numCPUs = 4
// )

// func Benchmark(b *testing.B) {
// 	b.StopTimer()
// 	b.ReportAllocs()
// 	runtime.GOMAXPROCS(numCPUs)

// 	schnorr := NewSchnorr(crypto.S256(), sha256.New())
// 	pri, _ := schnorr.KeyGen()
// 	b.StartTimer()
// 	b.RunParallel(func(pb *testing.PB) {
// 		for pb.Next() {
// 			schnorr.Sign(msg, pri)
// 		}
// 	})
// }

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	// runtime.GOMAXPROCS(numCPUs)
	// schnorr := NewSchnorr(S256(), sha256.New())
	schnorr := NewSchnorr(nil, nil)
	pri, _ := schnorr.KeyGen()
	for i := 0; i < b.N; i++ {
		schnorr.Sign(pri, msg)
	}
}
