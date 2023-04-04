package main

import (
	"blockchain-crypto/hash/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("abc")

	h := sha256.New()
	h.Write(input)
	bytes := h.Sum(nil)

	fmt.Println(hex.EncodeToString(bytes))
}
