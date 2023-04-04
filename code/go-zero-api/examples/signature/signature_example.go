package main

import (
	"blockchain-crypto/signature/bls"
	blsmulti "blockchain-crypto/signature/bls_multi"
	"blockchain-crypto/signature/bls_purego"
	"blockchain-crypto/signature/bls_threshold"
	ECDSA "blockchain-crypto/signature/ecdsa"
	ECSchnorr "blockchain-crypto/signature/ecschnorr"
	"crypto/sha256"
	"fmt"
)

func Example_ECDSA() {
	var seck, pubk, sig, mes string
	mes = "this is a test"
	h := sha256.New()
	h.Write([]byte(mes))
	meshashed := string(h.Sum(nil))
	seck, pubk = ECDSA.Keygen_api()
	sig = ECDSA.Sign_api(seck, meshashed)
	result := ECDSA.Verify_api(pubk, meshashed, sig)
	fmt.Println(result)
	//Output: true
}

func Example_ECSchnorr() {
	var seck, pubk, sigr, sigs, mes string
	mes = "this is a test"
	seck, pubk = ECSchnorr.Keygen_api()
	sigr, sigs = ECSchnorr.Sign_api(seck, mes)
	result := ECSchnorr.Verify_api(pubk, mes, sigr, sigs)
	fmt.Println(result)
	//Output: true
}

func Example_BLS() {
	var seck, pubk, sig, mes string
	mes = "this is a test"
	seck, pubk = bls.Keygen_api()
	sig = bls.Sign_api(seck, mes)
	result := bls.Verify_api(pubk, mes, sig)
	fmt.Println(result)
	//Output: true
}

func Example_BLS_purego() {
	var seck, pubk, sig, mes string
	mes = "this is a test"
	seck, pubk = bls_purego.Keygen_api()
	sig = bls_purego.Sign_api(seck, mes)
	result := bls_purego.Verify_api(pubk, mes, sig)
	fmt.Println(result)
	//Output: true
}

func Example_BLS_Multi() {
	var seck, pubk, mulsig, mes string
	var pubks, sigs []string
	var n int
	n = 5
	pubks = make([]string, n)
	sigs = make([]string, n)
	mes = "this is a test"
	for i := 0; i != n; i++ {
		seck, pubk = blsmulti.Keygen_api()
		sig := blsmulti.Sign_api(seck, mes)
		pubks[i] = pubk
		sigs[i] = sig
	}
	mulsig = blsmulti.MultiSign_api(sigs)
	result := blsmulti.MultiVerify_api(pubks, mes, mulsig)
	fmt.Println(result)
	//Output: true
}

func Example_BLS_Threshold() {
	var seck, pubk, sig, mes string
	var ids, secks, pubks []string
	var n, k, t int
	var result bool
	result = true
	n = 5
	t = 3
	mes = "this is a test"
	seck, pubk = bls.Keygen_api()
	sig = bls.Sign_api(seck, mes)
	ids, secks, pubks = bls_threshold.SecKeyShare_api(seck, t, n)

	k = 2 //k<t，无法正常生成
	ids2 := make([]string, k)
	secks2 := make([]string, k)
	pubks2 := make([]string, k)
	sigs2 := make([]string, k)
	for i := 0; i != k; i++ {
		ids2[i] = ids[i]
		secks2[i] = secks[i]
		pubks2[i] = pubks[i]
		sigs2[i] = bls.Sign_api(secks2[i], mes)
	}
	_, skr2Str := bls_threshold.SecKeyRecover_api(ids2, secks2)
	_, pkr2Str := bls_threshold.PubKeyRecover_api(ids2, pubks2)
	_, sigr2Str := bls_threshold.SigRecover_api(ids2, sigs2)
	if (skr2Str == seck) || (pkr2Str == pubk) || (sigr2Str == sig) {
		result = false
	}

	k = 4 //k>=t，可以正常生成
	ids4 := make([]string, k)
	secks4 := make([]string, k)
	pubks4 := make([]string, k)
	sigs4 := make([]string, k)
	for i := 0; i != k; i++ {
		ids4[i] = ids[i]
		secks4[i] = secks[i]
		pubks4[i] = pubks[i]
		sigs4[i] = bls.Sign_api(secks4[i], mes)
	}
	_, skr4Str := bls_threshold.SecKeyRecover_api(ids4, secks4)
	_, pkr4Str := bls_threshold.PubKeyRecover_api(ids4, pubks4)
	_, sigr4Str := bls_threshold.SigRecover_api(ids4, sigs4)
	if (skr4Str != seck) || (pkr4Str != pubk) || (sigr4Str != sig) {
		result = false
	}

	fmt.Println(result)
	//Output: true
}

func main() {
	/*Example_ECDSA()
	Example_ECSchnorr()
	Example_BLS()
	Example_EdDSA()
	Example_SM2()
	Example_BLS_Multi()
	Example_BLS_Threshold()*/
	Example_BLS_purego()
}
