package number

type Converter interface {
	GetTextWidth() int
	GetBitWidth() int
	GetBase() int

	StringToBytes(txt string) []byte
	BytesToString(bytes []byte) string
	IsStable(txt string) bool

	InputAcceptor() func(string, rune) bool
}
