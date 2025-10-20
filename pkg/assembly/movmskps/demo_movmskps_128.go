package movmskps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movmskps_128.s
var assemblyMovmskps128 string

//go:embed stub_movmskps_128.go
var stubMovmskps128 string

type MOVMSKPS128 struct {
}

func (v *MOVMSKPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *MOVMSKPS128) Output() *number.Parameter {
	return number.NewUintParameter(32, 32, 2)
}

func (v *MOVMSKPS128) Name() string {
	return "MOVMSKPS (128 bit) "
}

func (v *MOVMSKPS128) Description() string {
	return "TODO"
}

func (v *MOVMSKPS128) Stub() string {
	return stubMovmskps128
}

func (v *MOVMSKPS128) Assembly() string {
	return assemblyMovmskps128
}

func (v *MOVMSKPS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]byte{}

	movmskps128(&floats, &ret)

	log.Printf("MOVMSKPS128 input %v output %v", floats, ret)

	return ret[:]
}

func (v *MOVMSKPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
