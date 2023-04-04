// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"blockchain-crypto/hash/scrypt"
)

func main() {
	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length.
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}

	dk, err := scrypt.Key([]byte("some password"), salt, 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(dk))
	// Base64 Output: lGnMz8io0AUkfzn6Pls1qX20Vs7PGN6sbYQ2TQgY12M=
	//hex output: 9469cccfc8a8d005247f39fa3e5b35a97db456cecf18deac6d84364d0818d763
}
