package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_threes.s
var assemblyVshufps128Threes string

//go:embed stub_vshufps_128_threes.go
var stubVshufps128Threes string

type VSHUFPS128THREES struct {
}

func (v *VSHUFPS128THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSHUFPS128THREES) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSHUFPS128THREES) Name() string {
	return "VSHUFPS (128 bit) threes"
}

func (v *VSHUFPS128THREES) Description() string {
	return "VSHUFPS imm8=0xFF: dst = [a3,a3,b3,b3]"
}

func (v *VSHUFPS128THREES) Stub() string {
	return stubVshufps128Threes
}

func (v *VSHUFPS128THREES) Assembly() string {
	return assemblyVshufps128Threes
}

func (v *VSHUFPS128THREES) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vshufps128Threes(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS128THREES input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS128THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
