package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movups_128.s
var assemblyMovups128 string

//go:embed stub_movups_128.go
var stubMovups128 string

type MOVUPS128 struct {
}

func (v *MOVUPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *MOVUPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MOVUPS128) Name() string {
	return "MOVUPS XMM (128 bit)"
}

func (v *MOVUPS128) Description() string {
	return "TODO"
}

func (v *MOVUPS128) Stub() string {
	return stubMovups128
}

func (v *MOVUPS128) Assembly() string {
	return assemblyMovups128
}

func (v *MOVUPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	movups128(&floats, &ret)

	log.Printf("MOVUPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MOVUPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
