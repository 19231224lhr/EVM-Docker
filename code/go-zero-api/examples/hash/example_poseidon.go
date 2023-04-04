package main

import (
	"blockchain-crypto/hash/poseidon"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("abc")

	res, _ := poseidon.HashBytes(input)
	bytes := res.Bytes()
	fmt.Println(hex.EncodeToString(bytes))
	//	0101f67118a4df881bcb37724cab0a3d377a0d60fb72bb3566014850ade42deb
}
