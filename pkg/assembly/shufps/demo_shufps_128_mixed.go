package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_mixed.s
var assemblyShufps128Mixed string

//go:embed stub_shufps_128_mixed.go
var stubShufps128Mixed string

type SHUFPS128MIXED struct {
}

func (v *SHUFPS128MIXED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *SHUFPS128MIXED) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *SHUFPS128MIXED) Name() string {
	return "SHUFPS (128 bit) mixed"
}

func (v *SHUFPS128MIXED) Description() string {
	return "SHUFPS imm8=0xE4: dst = [a0,a1,b2,b3]"
}

func (v *SHUFPS128MIXED) Stub() string {
	return stubShufps128Mixed
}

func (v *SHUFPS128MIXED) Assembly() string {
	return assemblyShufps128Mixed
}

func (v *SHUFPS128MIXED) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	shufps128Mixed(&floats1, &floats2, &ret)

	log.Printf("SHUFPS128MIXED input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *SHUFPS128MIXED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
