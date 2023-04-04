package vrf

import (
	"bytes"

	"github.com/ProtonMail/go-ecvrf/ecvrf"
)

func Newprivatekey() (sk *ecvrf.PrivateKey, err error) {

	return ecvrf.GenerateKey(nil)
}
func GeneratePublickey(sk *ecvrf.PrivateKey) (*ecvrf.PublicKey, error) {

	return sk.Public()
}
func Prove(sk *ecvrf.PrivateKey, message []byte) (vrf0, proof []byte, err error) {

	return sk.Prove(message)
}
func Verify(pk *ecvrf.PublicKey, message, vrf0,proof []byte) (verified bool, err error) {

	r,vrf1,err:=pk.Verify(message, proof)
	verified= r && bytes.Equal(vrf0,vrf1)
	return verified,err
}

