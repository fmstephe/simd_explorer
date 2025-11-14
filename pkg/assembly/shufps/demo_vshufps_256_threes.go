package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_threes.s
var assemblyVshufps256Threes string

//go:embed stub_vshufps_256_threes.go
var stubVshufps256Threes string

type VSHUFPS256THREES struct {
}

func (v *VSHUFPS256THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSHUFPS256THREES) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSHUFPS256THREES) Name() string {
	return "VSHUFPS (256 bit) threes"
}

func (v *VSHUFPS256THREES) Description() string {
	return "VSHUFPS imm8=0xFF per 128-bit lane: dst = [a3,a3,b3,b3 | a7,a7,b7,b7]"
}

func (v *VSHUFPS256THREES) Stub() string {
	return stubVshufps256Threes
}

func (v *VSHUFPS256THREES) Assembly() string {
	return assemblyVshufps256Threes
}

func (v *VSHUFPS256THREES) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vshufps256Threes(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS256THREES input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS256THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
