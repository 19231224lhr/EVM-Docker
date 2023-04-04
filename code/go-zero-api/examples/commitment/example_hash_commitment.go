package main

import (
	"fmt"

	"blockchain-crypto/commitment/hash_commitment"
)

func main() {

	s := hash_commitment.Commit("hello", "19900")

	r := hash_commitment.Verify("hello", "19900", s)
	fmt.Println(r)
}
