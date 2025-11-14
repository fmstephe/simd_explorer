package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_zeros.s
var assemblyShufps128Zeros string

//go:embed stub_shufps_128_zeros.go
var stubShufps128Zeros string

type SHUFPS128ZEROS struct {
}

func (v *SHUFPS128ZEROS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SHUFPS128ZEROS) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SHUFPS128ZEROS) Name() string {
	return "SHUFPS (128 bit) zeros"
}

func (v *SHUFPS128ZEROS) Description() string {
	return "SHUFPS imm8=0x00: dst = [a0,a0,b0,b0]"
}

func (v *SHUFPS128ZEROS) Stub() string {
	return stubShufps128Zeros
}

func (v *SHUFPS128ZEROS) Assembly() string {
	return assemblyShufps128Zeros
}

func (v *SHUFPS128ZEROS) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	shufps128Zeros(&floats1, &floats2, &ret)

	log.Printf("SHUFPS128ZEROS input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SHUFPS128ZEROS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
