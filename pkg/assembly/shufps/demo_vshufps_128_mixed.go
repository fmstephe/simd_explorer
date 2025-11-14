package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_mixed.s
var assemblyVshufps128Mixed string

//go:embed stub_vshufps_128_mixed.go
var stubVshufps128Mixed string

type VSHUFPS128MIXED struct {
}

func (v *VSHUFPS128MIXED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSHUFPS128MIXED) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSHUFPS128MIXED) Name() string {
	return "VSHUFPS (128 bit) mixed"
}

func (v *VSHUFPS128MIXED) Description() string {
	return "VSHUFPS imm8=0xE4: dst = [a0,a1,b2,b3]"
}

func (v *VSHUFPS128MIXED) Stub() string {
	return stubVshufps128Mixed
}

func (v *VSHUFPS128MIXED) Assembly() string {
	return assemblyVshufps128Mixed
}

func (v *VSHUFPS128MIXED) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vshufps128Mixed(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS128MIXED input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS128MIXED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
