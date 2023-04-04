package main

import (
	"blockchain-crypto/hash/sha512"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("abc")

	h := sha512.New()
	h.Write(input)
	bytes := h.Sum(nil)

	fmt.Println(hex.EncodeToString(bytes))
	//	ddaf35a193617abacc417349ae20413112e6fa4e89a97ea20a9eeee64b55d39a2192992a274fc1a836ba3c23a3feebbd454d4423643ce80e2a9ac94fa54ca49f
}
