package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_reverse.s
var assemblyShufps128Reverse string

//go:embed stub_shufps_128_reverse.go
var stubShufps128Reverse string

type SHUFPS128REVERSE struct {
}

func (v *SHUFPS128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SHUFPS128REVERSE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SHUFPS128REVERSE) Name() string {
	return "SHUFPS (128 bit) reverse"
}

func (v *SHUFPS128REVERSE) Description() string {
	return "SHUFPS imm8=0x1B: dst = [a3,a2,b1,b0]"
}

func (v *SHUFPS128REVERSE) Stub() string {
	return stubShufps128Reverse
}

func (v *SHUFPS128REVERSE) Assembly() string {
	return assemblyShufps128Reverse
}

func (v *SHUFPS128REVERSE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	shufps128Reverse(&floats1, &floats2, &ret)

	log.Printf("SHUFPS128REVERSE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SHUFPS128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
