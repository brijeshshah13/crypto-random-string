package cryptorandomstring

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
)

func (g *Generator) Generate() string {
	return generateRandomBytes(uint64(math.Ceil(float64(g.length)*0.5)), "hex", g.length)
}

func Generate() string {
	return defaultGenerator.Generate()
}

func generateRandomBytes(byteLength uint64, kind string, length uint64) string {
	b := make([]byte, byteLength)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	test := hex.EncodeToString(b)
	fmt.Println(test[0:length])
	return test[0:length]
}
