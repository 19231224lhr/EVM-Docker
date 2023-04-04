package main

import (
	"blockchain-crypto/vdf/wesolowski_go"
	"encoding/base64"

	// "crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func main() {
	input := [32]byte{0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe,
		0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef, 0xde, 0xad, 0xbe, 0xef}

	// json
	output_string := base64.StdEncoding.EncodeToString(input[:])
	fmt.Println(output_string)

	vdf := wesolowski_go.New(100, input)

	outputChannel := vdf.GetOutputChannel()
	start := time.Now()

	vdf.Execute()

	duration := time.Now().Sub(start)

	output := <-outputChannel
	fmt.Println("===", input)
	fmt.Println("===", output)

	log.Println(fmt.Sprintf("VDF computation finished, result is  %s", hex.EncodeToString(output[:])))
	log.Println(fmt.Sprintf("VDF computation finished, time spent %s", duration.String()))
	log.Println(fmt.Sprintf("VDF verify %t", vdf.Verify(output)))
}
