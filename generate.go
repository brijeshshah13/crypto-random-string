package cryptorandomstring

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
)

func (g *Generator) Generate() (string, error) {
	// 1: check if length >= 0 && length is a finite number, handled as the data type is uint64
	// 2: check if both, kind and characters are passed
	if g.kind != "" && g.characters != "" {
		return "", fmt.Errorf("expected either `kind` or `characters`")
	}
	// 3: check if characters is passed and is of type string, handled as the data type is string
	// 4: check if kind is one of the allowed types
	if !allowedTypes[g.kind] {
		return "", fmt.Errorf("unknown type: %s", g.kind)
	}
	// 5: check if both, kind and characters are not passed
	if g.kind == "" && g.characters == "" {
		g.WithKind(defaultKind)
	}
	// 6: check if kind is equal to hex
	if g.kind == "hex" {
		if str, err := generateRandomBytes(uint64(math.Ceil(float64(g.length)*0.5)), "hex", g.length); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 7: check if kind is equal to base64
	if g.kind == "base64" {
		if str, err := generateRandomBytes(uint64(math.Ceil(float64(g.length)*0.75)), "base64", g.length); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 8: check if kind is equal to url-safe
	if g.kind == "url-safe" {

	}
	// 9: check if kind is equal to numeric
	if g.kind == "numeric" {

	}
	// 10: check if kind is equal to distinguishable
	if g.kind == "distinguishable" {

	}
	// 11: check if kind is equal to ascii-printable
	if g.kind == "ascii-printable" {

	}
	// 12: check if kind is equal to alphanumeric
	if g.kind == "alphanumeric" {

	}
	// 13: check if length of characters is 0, handled by #5
	// 14: check if length of characters is greater than 65536
	if len(g.characters) > 0x10000 {
		return "", fmt.Errorf("expected `characters` string length to be less or equal to 65536")
	}
	return "test", nil
}

func Generate() (string, error) {
	return defaultGenerator.Generate()
}

func generateRandomBytes(byteLength uint64, kind string, length uint64) (string, error) {
	randomBytes := make([]byte, byteLength)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	if kind == "hex" {
		return hex.EncodeToString(randomBytes)[0:length], nil
	} else if kind == "base64" {
		return base64.StdEncoding.EncodeToString(randomBytes)[0:length], nil
	} else {
		return "", nil
	}
}
