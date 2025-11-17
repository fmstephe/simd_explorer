package unpckhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpckhps_256.s
var assemblyVunpckhps256 string

//go:embed stub_vunpckhps_256.go
var stubVunpckhps256 string

type VUNPCKHPS256 struct {
}

func (v *VUNPCKHPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VUNPCKHPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VUNPCKHPS256) Name() string {
	return "VUNPCKHPS (256 bit) interleave-high"
}

func (v *VUNPCKHPS256) Description() string {
	return "Unpack and interleave high 64-bit halves per 128-bit lane: [a2,b2,a3,b3 | a6,b6,a7,b7]."
}

func (v *VUNPCKHPS256) Stub() string {
	return stubVunpckhps256
}

func (v *VUNPCKHPS256) Assembly() string {
	return assemblyVunpckhps256
}

func (v *VUNPCKHPS256) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vunpckhps256(&floats1, &floats2, &ret)

	log.Printf("VUNPCKHPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VUNPCKHPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
