package main

import (
	"blockchain-crypto/bulletproofs"
	"fmt"
	"math/big"
)

func main() {
	v := big.NewInt(13) // the scalar we want to generate a range proof for
	gamma := big.NewInt(10)
	prover := bulletproofs.NewProver(4)

	// V = γH + vG.
	V := bulletproofs.Commit(gamma, prover.BlindingGenerator, v, prover.ValueGenerator)

	proof, err := prover.CreateRangeProof(V, v, gamma, [32]byte{}, [16]byte{})
	if err != nil {
		fmt.Printf("failed to create range proof: %v\n", err)
	}
	fmt.Println("Range proof is created")

	if !prover.Verify(V, proof) {
		fmt.Printf("Expected valid proof\n")
	}
	fmt.Println("Range proof is valid")
}
