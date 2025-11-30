package assembly

import "github.com/fmstephe/simd_explorer/pkg/ui/number"

type Instruction interface {
	Inputs() []*number.Parameter
	Output() *number.Parameter
	Name() string
	Description() string
	Stub() string
	Assembly() string
	Run() []byte
	Supported() bool
}
