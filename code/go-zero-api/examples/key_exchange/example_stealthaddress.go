package main

import (
	"fmt"
	"math/big"

	"blockchain-crypto/key_exchange/stealth_address"
)

func main() {

	a, A, b, B := stealth_address.RecCalculateKeyPairs("hello", "hi")

	P := stealth_address.SendCalculateObfuscateAddress(new(big.Int).SetInt64(1000), A, B)

	R := stealth_address.SendCalculatePublicKey(new(big.Int).SetInt64(1000))

	t := stealth_address.RecCalculateObfuscateAddress(P, R, a, B)
	fmt.Println(t)

	priv := stealth_address.RecCalculateAddressPrivatekey(a, b, R)
	fmt.Println(priv)
}
