package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movaps_128.s
var assemblyMovaps128 string

//go:embed stub_movaps_128.go
var stubMovaps128 string

type MOVAPS128 struct {
}

func (v *MOVAPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *MOVAPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVAPS128) Name() string {
	return "MOVAPS XMM (128 bit)"
}

func (v *MOVAPS128) Description() string {
	return "Aligned move of packed single-precision floats between memory and XMM; copies data unchanged."
}

func (v *MOVAPS128) Stub() string {
	return stubMovaps128
}

func (v *MOVAPS128) Assembly() string {
	return assemblyMovaps128
}

func (v *MOVAPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movaps128(&floats, &ret)

	log.Printf("MOVAPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVAPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
