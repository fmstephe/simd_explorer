package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_ones.s
var assemblyShufps128Ones string

//go:embed stub_shufps_128_ones.go
var stubShufps128Ones string

type SHUFPS128ONES struct {
}

func (v *SHUFPS128ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SHUFPS128ONES) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SHUFPS128ONES) Name() string {
	return "SHUFPS (128 bit) ones"
}

func (v *SHUFPS128ONES) Description() string {
	return "SHUFPS imm8=0x55: dst = [a1,a1,b1,b1]"
}

func (v *SHUFPS128ONES) Stub() string {
	return stubShufps128Ones
}

func (v *SHUFPS128ONES) Assembly() string {
	return assemblyShufps128Ones
}

func (v *SHUFPS128ONES) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	shufps128Ones(&floats1, &floats2, &ret)

	log.Printf("SHUFPS128ONES input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SHUFPS128ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
