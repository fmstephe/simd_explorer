package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_twos.s
var assemblyVshufps256Twos string

//go:embed stub_vshufps_256_twos.go
var stubVshufps256Twos string

type VSHUFPS256TWOS struct {
}

func (v *VSHUFPS256TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSHUFPS256TWOS) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSHUFPS256TWOS) Name() string {
	return "VSHUFPS (256 bit) twos"
}

func (v *VSHUFPS256TWOS) Description() string {
	return "VSHUFPS imm8=0xAA per 128-bit lane: dst = [a2,a2,b2,b2 | a6,a6,b6,b6]"
}

func (v *VSHUFPS256TWOS) Stub() string {
	return stubVshufps256Twos
}

func (v *VSHUFPS256TWOS) Assembly() string {
	return assemblyVshufps256Twos
}

func (v *VSHUFPS256TWOS) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vshufps256Twos(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS256TWOS input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS256TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
