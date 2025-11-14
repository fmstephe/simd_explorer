package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_reverse.s
var assemblyVshufps256Reverse string

//go:embed stub_vshufps_256_reverse.go
var stubVshufps256Reverse string

type VSHUFPS256REVERSE struct {
}

func (v *VSHUFPS256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSHUFPS256REVERSE) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSHUFPS256REVERSE) Name() string {
	return "VSHUFPS (256 bit) reverse"
}

func (v *VSHUFPS256REVERSE) Description() string {
	return "VSHUFPS imm8=0x1B per 128-bit lane: dst = [a3,a2,b1,b0 | a7,a6,b5,b4]"
}

func (v *VSHUFPS256REVERSE) Stub() string {
	return stubVshufps256Reverse
}

func (v *VSHUFPS256REVERSE) Assembly() string {
	return assemblyVshufps256Reverse
}

func (v *VSHUFPS256REVERSE) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vshufps256Reverse(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS256REVERSE input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
