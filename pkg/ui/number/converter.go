package number

type Converter interface {
	GetTextWidth() int
	GetBitWidth() int

	StringToBytes(txt string) []byte
	BytesToString(bytes []byte) string

	InputAcceptor() func(string, rune) bool
}
