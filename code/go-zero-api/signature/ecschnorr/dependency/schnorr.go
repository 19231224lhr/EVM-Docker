package dependency

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"hash"
	"math/big"
)

var (
	ErrSigLen   = errors.New("Signature Length Invalid, Must be 64")
	ErrSigCheck = errors.New("Not a valid ec-schnorr signature")
)

// combinedMult implements fast multiplication S1*g + S2*p (g - generator, p - arbitrary point)
type combinedMult interface {
	CombinedMult(bigX, bigY *big.Int, baseScalar, scalar []byte) (x, y *big.Int)
}

type Schnorr struct {
	curve elliptic.Curve
	h     hash.Hash
}

func NewSchnorr(curve elliptic.Curve, h hash.Hash) *Schnorr {
	if curve == nil {
		curve = S256()
	}
	if h == nil {
		h = sha256.New()
	}

	return &Schnorr{curve, h}
}

// signECShnorr
func (schnorr *Schnorr) Sign(prv *ecdsa.PrivateKey, msg []byte) (r, s *big.Int, err error) {
	h := schnorr.h
	//random k
	params := prv.Params()
	//bigOne := new(big.Int).SetUint64(1)
	k, err := rand.Int(rand.Reader, params.N)
	if err != nil {
		return nil, nil, fmt.Errorf("Get rand k failed [%v]", err)
	}

	//compute r=H(R, pk, m)
	N := (params.BitSize + 7) / 8
	x, y := prv.Curve.ScalarBaseMult(k.Bytes())
	//buf := point2oct(prv.Curve.ScalarBaseMult(k.Bytes()), params.BitSize/8)
	buf := point2oct(x, y, N)
	if _, err = h.Write(buf); err != nil {
		return
	}

	buf = point2oct(prv.PublicKey.X, prv.PublicKey.Y, N)
	if _, err = h.Write(buf); err != nil {
		return
	}
	//hash the message
	if _, err = h.Write(msg); err != nil {
		return
	}
	e := h.Sum(nil)
	//to big.Int, nmod
	r = new(big.Int).SetBytes(e)
	r = r.Mod(r, params.N)
	h.Reset()

	//compute s = k-dr
	dr := mulMod(prv.D, r, params.N)
	s = subMod(k, dr, params.N)

	//fmt.Println(len(s.Bytes()), len(r.Bytes()))
	return r, s, nil
}

// verifyECShnorr
func (schnorr *Schnorr) Verify(pub *ecdsa.PublicKey, r, s *big.Int, msg []byte) (valid bool, err error) {
	h := schnorr.h
	// var signature ecdsaSignature
	// _, err = asn1.Unmarshal(sig, &signature)
	// if err != nil {
	// 	return false, err
	// }
	// s, r := signature.R, signature.S
	params := pub.Params()
	bigOne := new(big.Int).SetUint64(1)
	if s.Cmp(params.N.Sub(params.N, bigOne)) > 0 || r.Cmp(params.N.Sub(params.N, bigOne)) > 0 {
		return false, errors.New("s, r is not valid")
	}

	//compute Q = sG+r*pk
	var grx, grv *big.Int
	if opt, ok := pub.Curve.(combinedMult); ok {
		grx, grv = opt.CombinedMult(pub.X, pub.Y, s.Bytes(), r.Bytes())
	} else {
		gsx, gsy := pub.ScalarBaseMult(s.Bytes())
		gex, gey := pub.ScalarMult(pub.X, pub.Y, r.Bytes())
		grx, grv = pub.Add(gsx, gsy, gex, gey)
	}

	N := (params.BitSize + 7) / 8
	buf := point2oct(grx, grv, N)
	if _, err = h.Write(buf); err != nil {
		return
	}

	buf = point2oct(pub.X, pub.Y, N)
	if _, err = h.Write(buf); err != nil {
		return
	}
	//hash the message
	if _, err = h.Write(msg); err != nil {
		return
	}
	e := h.Sum(nil)

	rr := new(big.Int).SetBytes(e)
	rr = rr.Mod(rr, params.N)
	h.Reset()

	//fmt.Printf("r: %v\nrr: %v\n", r, rr)

	if r.Cmp(rr) != 0 {
		return false, ErrSigCheck
	}

	return true, nil
}

// func (ec *Schnorr) KeyGen() (privateKey, publicKey []byte, err error) {
func (ec *Schnorr) KeyGen() (pri *ecdsa.PrivateKey, err error) {
	pri, err = ecdsa.GenerateKey(ec.curve, rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("ecdsa key generate failed for [%v][%s]", ec.curve, err)
	}
	// privateKey = pri.D.Bytes()
	// publicKey = elliptic.Marshal(ec.curve, pri.X, pri.Y) //prefix 0x04
	return
}
