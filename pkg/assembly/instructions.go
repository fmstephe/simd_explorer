package assembly

type Instruction interface {
	InputSizes() []int
	OutputSize() int
	Name() string
	Description() string
	Stub() string
	Assembly() string
	Run([][]byte) []byte
	Supported() bool
}
