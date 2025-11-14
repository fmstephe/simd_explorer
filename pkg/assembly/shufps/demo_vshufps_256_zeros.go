package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_zeros.s
var assemblyVshufps256Zeros string

//go:embed stub_vshufps_256_zeros.go
var stubVshufps256Zeros string

type VSHUFPS256ZEROS struct {
}

func (v *VSHUFPS256ZEROS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSHUFPS256ZEROS) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSHUFPS256ZEROS) Name() string {
	return "VSHUFPS (256 bit) zeros"
}

func (v *VSHUFPS256ZEROS) Description() string {
	return "VSHUFPS imm8=0x00 per 128-bit lane: dst = [a0,a0,b0,b0 | a4,a4,b4,b4]"
}

func (v *VSHUFPS256ZEROS) Stub() string {
	return stubVshufps256Zeros
}

func (v *VSHUFPS256ZEROS) Assembly() string {
	return assemblyVshufps256Zeros
}

func (v *VSHUFPS256ZEROS) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vshufps256Zeros(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS256ZEROS input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS256ZEROS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
