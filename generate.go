package cryptorandomstring

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
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
		if str, err := generateForCustomCharacters(g.length, urlSafeCharacters); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 9: check if kind is equal to numeric
	if g.kind == "numeric" {
		if str, err := generateForCustomCharacters(g.length, numericCharacters); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 10: check if kind is equal to distinguishable
	if g.kind == "distinguishable" {
		if str, err := generateForCustomCharacters(g.length, distinguishableCharacters); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 11: check if kind is equal to ascii-printable
	if g.kind == "ascii-printable" {
		if str, err := generateForCustomCharacters(g.length, asciiPrintableCharacters); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 12: check if kind is equal to alphanumeric
	if g.kind == "alphanumeric" {
		if str, err := generateForCustomCharacters(g.length, alphanumericCharacters); err != nil {
			return "", err
		} else {
			return str, nil
		}
	}
	// 13: check if length of characters is 0, handled by #5
	// 14: check if length of characters is greater than 65536
	if len(g.characters) > 0x10000 {
		return "", fmt.Errorf("expected `characters` string length to be less or equal to 65536")
	}
	if str, err := generateForCustomCharacters(g.length, strings.Split(g.characters, "")); err != nil {
		return "", err
	} else {
		return str, nil
	}
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

func generateForCustomCharacters(length uint64, characters []string) (string, error) {
	// Generating entropy is faster than complex math operations, so we use the simplest way
	characterCount := len(characters)
	maxValidSelector := uint16((math.Floor(0x10000/float64(characterCount)) * float64(characterCount)) - 1) // Using values above this will ruin distribution when using modular division
	entropyLength := int(2 * math.Ceil(1.1*float64(length)))                                                // Generating a bit more than required so chances we need more than one pass will be really low
	var generatedString string
	var generatedStringLength uint64

	for generatedStringLength < length {
		randomBytes := make([]byte, entropyLength)
		if _, err := rand.Read(randomBytes); err != nil {
			return "", err
		}
		entropyPosition := 0

		for entropyPosition < entropyLength && generatedStringLength < length {
			entropyValue := binary.LittleEndian.Uint16(randomBytes[entropyPosition:])
			entropyPosition += 2
			if entropyValue > maxValidSelector { // Skip values which will ruin distribution when using modular division
				continue
			}

			generatedString += characters[entropyValue%uint16(characterCount)]
			generatedStringLength++
		}
	}
	return generatedString, nil
}
