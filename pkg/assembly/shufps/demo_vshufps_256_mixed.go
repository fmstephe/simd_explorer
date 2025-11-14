package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_mixed.s
var assemblyVshufps256Mixed string

//go:embed stub_vshufps_256_mixed.go
var stubVshufps256Mixed string

type VSHUFPS256MIXED struct {
}

func (v *VSHUFPS256MIXED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSHUFPS256MIXED) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSHUFPS256MIXED) Name() string {
	return "VSHUFPS (256 bit) mixed"
}

func (v *VSHUFPS256MIXED) Description() string {
	return "VSHUFPS imm8=0xE4 per 128-bit lane: dst = [a0,a1,b2,b3 | a4,a5,b6,b7]"
}

func (v *VSHUFPS256MIXED) Stub() string {
	return stubVshufps256Mixed
}

func (v *VSHUFPS256MIXED) Assembly() string {
	return assemblyVshufps256Mixed
}

func (v *VSHUFPS256MIXED) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vshufps256Mixed(&floats1, &floats2, &ret)

	log.Printf("VSHUFPS256MIXED input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSHUFPS256MIXED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
