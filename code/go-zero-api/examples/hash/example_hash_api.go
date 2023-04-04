package main

import (
	"blockchain-crypto/hash"
	"encoding/hex"
	"fmt"
)

func main() {
	input := []byte("test")

	// sha2-256
	output1, err1 := hash.Hash("sha256", input)
	if err1 != nil {
		fmt.Println("sha256: ", err1)
	}
	fmt.Println("sha256:\n" + hex.EncodeToString(output1))

	// sha2-512
	output2, err2 := hash.Hash("sha512", input)
	if err2 != nil {
		fmt.Println("sha512: ", err2)
	}
	fmt.Println("sha512:\n" + hex.EncodeToString(output2))

	// sha3-256
	output3, err3 := hash.Hash("sha3-256", input)
	if err3 != nil {
		fmt.Println("sha3-256: ", err3)
	}
	fmt.Println("sha3-256:\n" + hex.EncodeToString(output3))

	// sha3-512
	output4, err4 := hash.Hash("sha3-512", input)
	if err4 != nil {
		fmt.Println("sha3-512: ", err4)
	}
	fmt.Println("sha3-512:\n" + hex.EncodeToString(output4))

	// keccak-256
	output5, err5 := hash.Hash("keccak256", input)
	if err5 != nil {
		fmt.Println("keccak256: ", err5)
	}
	fmt.Println("keccak256:\n" + hex.EncodeToString(output5))

	// keccak-512
	output6, err6 := hash.Hash("keccak512", input)
	if err6 != nil {
		fmt.Println("keccak512: ", err6)
	}
	fmt.Println("keccak512:\n" + hex.EncodeToString(output6))

	// poseidon-256
	output7, err7 := hash.Hash("poseidon-256", input)
	if err7 != nil {
		fmt.Println("poseidon-256: ", err7)
	}
	fmt.Println("poseidon-256:\n" + hex.EncodeToString(output7))

	// scrypt
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	args := hash.WithArgs(salt, 1<<15, 8, 1, 32)
	output8, err8 := hash.Hash("scrypt", input, args)
	if err8 != nil {
		fmt.Println("scrypt: ", err8)
	}
	fmt.Println("scrypt:\n" + hex.EncodeToString(output8))

	// wrong type
	output9, err9 := hash.Hash("wrong_type", input)
	if err9 != nil {
		fmt.Println("wrong_type: " + err9.Error())
	}
	fmt.Println("wrong_type:\n" + hex.EncodeToString(output9))

	// ripemd160
	output10, err10 := hash.Hash("ripemd160", input)
	if err10 != nil {
		fmt.Println("sha256: ", err10)
	}
	fmt.Println("ripemd160:\n" + hex.EncodeToString(output10))
}
