package number

type Converter interface {
	GoType() string
	GetTextWidth() int
	GetBitWidth() int
	GetBase() int

	StringToBytes(txt string) []byte
	BytesToString(bytes []byte) string
	// It is critical that the normalised value returned here will not
	// change when passed into this method more than once. This is critical
	// because it is used in a recursive text-changed listener and the text
	// is recursively set if the normalised value is different from the
	// input value. If the normalised value changes when re-normalised an
	// infinite recursive loop will occur.
	Normalised(txt string) (string, bool)
	IsStable(txt string) bool

	InputAcceptor() func(string, rune) bool
}
