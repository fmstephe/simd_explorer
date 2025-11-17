package unpckhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpckhps_128.s
var assemblyVunpckhps128 string

//go:embed stub_vunpckhps_128.go
var stubVunpckhps128 string

type VUNPCKHPS128 struct {
}

func (v *VUNPCKHPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VUNPCKHPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VUNPCKHPS128) Name() string {
	return "VUNPCKHPS (128 bit) interleave-high"
}

func (v *VUNPCKHPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *VUNPCKHPS128) Stub() string {
	return stubVunpckhps128
}

func (v *VUNPCKHPS128) Assembly() string {
	return assemblyVunpckhps128
}

func (v *VUNPCKHPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vunpckhps128(&floats1, &floats2, &ret)

	log.Printf("VUNPCKHPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VUNPCKHPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
