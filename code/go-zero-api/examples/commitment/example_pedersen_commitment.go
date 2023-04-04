package main

import (
	"fmt"

	"blockchain-crypto/commitment/pedersen_commitment"
)

func main() {
	
	G, H := pedersen_commitment.ParamsGen()
	r := pedersen_commitment.RandomGen()
	s := uint64(100)
	comit := pedersen_commitment.Commit(G, H, s, r)
	res := pedersen_commitment.Verify(comit, G, H, s, r)
	fmt.Print(res)
}
