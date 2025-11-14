package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_zeros.s
var assemblyVshufps128Zeros string

//go:embed stub_vshufps_128_zeros.go
var stubVshufps128Zeros string

type VSHUFPS128ZEROS struct {
}

func (v *VSHUFPS128ZEROS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSHUFPS128ZEROS) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSHUFPS128ZEROS) Name() string {
	return "VSHUFPS (128 bit) zeros"
}

func (v *VSHUFPS128ZEROS) Description() string {
	return "VSHUFPS imm8=0x00: dst = [a0,a0,b0,b0]"
}

func (v *VSHUFPS128ZEROS) Stub() string {
	return stubVshufps128Zeros
}

func (v *VSHUFPS128ZEROS) Assembly() string {
	return assemblyVshufps128Zeros
}

func (v *VSHUFPS128ZEROS) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vshufps128Zeros(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS128ZEROS input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS128ZEROS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
