package main

import (
	"blockchain-crypto/vdf/wesolowski_go"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func RepeatedSquare(x *wesolowski_go.ClassGroup, k int) *wesolowski_go.ClassGroup {
	defer timeTrack(time.Now())

	for i := 0; i < k; i++ {
		x = x.Square()
	}

	return x
}

func RepeatedSquareSlow(x *wesolowski_go.ClassGroup, k int) *wesolowski_go.ClassGroup {
	defer timeTrack(time.Now())

	for i := 0; i < k; i++ {
		x = x.SquareUsingMultiply()
	}

	return x
}

func TestTwoSquarePerformance(t *testing.T) {
	for k := 0; k < 10; k++ {
		seed := make([]byte, 32)
		rand.Read(seed)

		D := wesolowski_go.CreateDiscriminant(seed, 2048)
		x := wesolowski_go.NewClassGroupFromAbDiscriminant(big.NewInt(2), big.NewInt(1), D)

		y := wesolowski_go.CloneClassGroup(x)
		y1 := wesolowski_go.CloneClassGroup(x)

		y = RepeatedSquare(y, 5000)
		y1 = RepeatedSquareSlow(y1, 5000)

		assert.Equal(t, true, y.Equal(y1), "k=%d, seed=%s", k, hex.EncodeToString(seed))
		log.Print(fmt.Sprintf("Test case %d good", k))
	}
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}
