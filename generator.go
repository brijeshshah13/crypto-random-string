package cryptorandomstring

type Generator struct {
	length     uint64
	kind       string
	characters string
}

const defaultLength = 10

const defaultKind = "hex"

var defaultGenerator *Generator

func New() *Generator {
	return &Generator{
		length: defaultLength,
		kind:   defaultKind,
	}
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

func init() {
	defaultGenerator = New()
}
