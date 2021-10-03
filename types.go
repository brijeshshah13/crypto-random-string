package cryptorandomstring

type (
	generateForCustomCharactersFunc func(length uint64, characters string) string
	generateRandomBytesFunc         func(byteLength uint64, kind string, length uint64) string
)
