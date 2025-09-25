package assembly

type Instruction interface {
	InputSize() int
	OutputSize() int
	Name() string
	Description() string
	Stub() string
	Assembly() string
	Run([]byte) []byte
	Supported() bool
}
