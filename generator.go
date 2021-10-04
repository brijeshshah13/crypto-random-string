package cryptorandomstring

import "strings"

type Generator struct {
	length     uint64
	kind       string
	characters string
}

const defaultKind = "hex"

var (
	urlSafeCharacters         = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~", "")
	numericCharacters         = strings.Split("0123456789", "")
	distinguishableCharacters = strings.Split("CDEHKMPRTUWXY012458", "")
	asciiPrintableCharacters  = strings.Split("!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", "")
	alphanumericCharacters    = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "")
)

var allowedTypes = map[string]bool{
	"":                true,
	"hex":             true,
	"base64":          true,
	"url-safe":        true,
	"numeric":         true,
	"distinguishable": true,
	"ascii-printable": true,
	"alphanumeric":    true,
}

var defaultGenerator *Generator

func New() *Generator {
	return &Generator{}
}

func (g *Generator) WithLength(length uint64) *Generator {
	g.length = length
	return g
}

func WithLength(length uint64) *Generator {
	return defaultGenerator.WithLength(length)
}

func (g *Generator) WithKind(kind string) *Generator {
	g.kind = kind
	return g
}

func WithKind(kind string) *Generator {
	return defaultGenerator.WithKind(kind)
}

func (g *Generator) WithCharacters(characters string) *Generator {
	g.characters = characters
	return g
}

func WithCharacters(characters string) *Generator {
	return defaultGenerator.WithCharacters(characters)
}

func init() {
	defaultGenerator = New()
}
