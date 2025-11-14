package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_reverse.s
var assemblyVshufps128Reverse string

//go:embed stub_vshufps_128_reverse.go
var stubVshufps128Reverse string

type VSHUFPS128REVERSE struct {
}

func (v *VSHUFPS128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSHUFPS128REVERSE) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSHUFPS128REVERSE) Name() string {
	return "VSHUFPS (128 bit) reverse"
}

func (v *VSHUFPS128REVERSE) Description() string {
	return "VSHUFPS imm8=0x1B: dst = [a3,a2,b1,b0]"
}

func (v *VSHUFPS128REVERSE) Stub() string {
	return stubVshufps128Reverse
}

func (v *VSHUFPS128REVERSE) Assembly() string {
	return assemblyVshufps128Reverse
}

func (v *VSHUFPS128REVERSE) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vshufps128Reverse(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS128REVERSE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
