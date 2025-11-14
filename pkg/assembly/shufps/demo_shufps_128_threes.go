package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_threes.s
var assemblyShufps128Threes string

//go:embed stub_shufps_128_threes.go
var stubShufps128Threes string

type SHUFPS128THREES struct {
}

func (v *SHUFPS128THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SHUFPS128THREES) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SHUFPS128THREES) Name() string {
	return "SHUFPS (128 bit) threes"
}

func (v *SHUFPS128THREES) Description() string {
	return "SHUFPS imm8=0xFF: dst = [a3,a3,b3,b3]"
}

func (v *SHUFPS128THREES) Stub() string {
	return stubShufps128Threes
}

func (v *SHUFPS128THREES) Assembly() string {
	return assemblyShufps128Threes
}

func (v *SHUFPS128THREES) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	shufps128Threes(&floats1, &floats2, &ret)

	log.Printf("SHUFPS128THREES input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SHUFPS128THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
