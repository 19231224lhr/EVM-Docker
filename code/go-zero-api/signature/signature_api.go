package signature

import (
	bls "blockchain-crypto/signature/bls"
	blsgo "blockchain-crypto/signature/bls_purego"
	ecdsa "blockchain-crypto/signature/ecdsa"
	ecschnorr "blockchain-crypto/signature/ecschnorr"
	eddsa "blockchain-crypto/signature/eddsa"
	sm2 "blockchain-crypto/signature/sm2"
)

func KeygenApi(scheme string) ([]string, []string) {
	switch scheme {
	case "bls":
		{
			seck := make([]string, 1)
			pubk := make([]string, 1)
			seck[0], pubk[0] = bls.Keygen_api()
			return seck, pubk
		}
	case "bls_purego":
		{
			seck := make([]string, 1)
			pubk := make([]string, 1)
			seck[0], pubk[0] = blsgo.Keygen_api()
			return seck, pubk
		}
	case "ecdsa":
		{
			seck := make([]string, 1)
			pubk := make([]string, 1)
			seck[0], pubk[0] = ecdsa.Keygen_api()
			return seck, pubk
		}
	case "ecschnorr":
		{
			seck := make([]string, 1)
			pubk := make([]string, 1)
			seck[0], pubk[0] = ecschnorr.Keygen_api()
			return seck, pubk
		}
	case "eddsa":
		{
			seck := make([]string, 1)
			pubk := make([]string, 1)
			seck[0], pubk[0] = eddsa.Keygen_api()
			return seck, pubk
		}
	case "sm2":
		{
			seck := make([]string, 1)
			pubk := make([]string, 1)
			seck[0], pubk[0] = sm2.Keygen_api()
			return seck, pubk
		}
	default:
		{
			println("Wrong scheme name.")
			return nil, nil
		}
	}
}

func SignApi(scheme string, seck []string, mes string) []string {
	switch scheme {
	case "bls":
		{
			sig := make([]string, 1)
			sig[0] = bls.Sign_api(seck[0], mes)
			return sig
		}
	case "bls_purego":
		{
			sig := make([]string, 1)
			sig[0] = blsgo.Sign_api(seck[0], mes)
			return sig
		}
	case "ecdsa":
		{
			sig := make([]string, 1)
			sig[0] = ecdsa.Sign_api(seck[0], mes)
			return sig
		}
	case "ecschnorr":
		{
			sig := make([]string, 2)
			sig[0], sig[1] = ecschnorr.Sign_api(seck[0], mes)
			return sig
		}
	case "eddsa":
		{
			sig := make([]string, 1)
			sig[0] = eddsa.Sign_api(seck[0], mes)
			return sig
		}
	case "sm2":
		{
			sig := make([]string, 1)
			sig[0] = sm2.Sign_api(seck[0], mes)
			return sig
		}
	default:
		{
			println("Wrong scheme name.")
			return nil
		}
	}
}

func VerifyApi(scheme string, pubk []string, mes string, sig []string) bool {
	var result bool
	switch scheme {
	case "bls":
		{
			result = bls.Verify_api(pubk[0], mes, sig[0])
		}
	case "bls_purego":
		{
			result = blsgo.Verify_api(pubk[0], mes, sig[0])
		}
	case "ecdsa":
		{
			result = ecdsa.Verify_api(pubk[0], mes, sig[0])
		}
	case "ecschnorr":
		{
			result = ecschnorr.Verify_api(pubk[0], mes, sig[0], sig[1])
		}
	case "eddsa":
		{
			result = eddsa.Verify_api(pubk[0], mes, sig[0])
		}
	case "sm2":
		{
			result = sm2.Verify_api(pubk[0], mes, sig[0])
		}
	default:
		{
			println("Wrong scheme name.")
		}
	}
	return result
}
