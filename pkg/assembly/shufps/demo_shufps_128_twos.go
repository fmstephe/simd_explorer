package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_twos.s
var assemblyShufps128Twos string

//go:embed stub_shufps_128_twos.go
var stubShufps128Twos string

type SHUFPS128TWOS struct {
}

func (v *SHUFPS128TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SHUFPS128TWOS) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SHUFPS128TWOS) Name() string {
	return "SHUFPS (128 bit) twos"
}

func (v *SHUFPS128TWOS) Description() string {
	return "SHUFPS imm8=0xAA: dst = [a2,a2,b2,b2]"
}

func (v *SHUFPS128TWOS) Stub() string {
	return stubShufps128Twos
}

func (v *SHUFPS128TWOS) Assembly() string {
	return assemblyShufps128Twos
}

func (v *SHUFPS128TWOS) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	shufps128Twos(&floats1, &floats2, &ret)

	log.Printf("SHUFPS128TWOS input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SHUFPS128TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
