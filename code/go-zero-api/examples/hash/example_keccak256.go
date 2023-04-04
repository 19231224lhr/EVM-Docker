package main

import (
	"blockchain-crypto/hash/sha3"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("abc")

	h := sha3.NewLegacyKeccak256()
	h.Write(input)
	bytes := h.Sum(nil)

	fmt.Println(hex.EncodeToString(bytes))
}
