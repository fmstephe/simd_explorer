package assembly

import "github.com/fmstephe/simd_explorer/pkg/ui/number"

type Instruction interface {
	InputSizes() []number.Converter
	OutputSize() number.Converter
	Name() string
	Description() string
	Stub() string
	Assembly() string
	Run([][]byte) []byte
	Supported() bool
}
