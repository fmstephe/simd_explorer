package unpckhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_unpckhps_128.s
var assemblyUnpckhps128 string

//go:embed stub_unpckhps_128.go
var stubUnpckhps128 string

type UNPCKHPS128 struct {
}

func (v *UNPCKHPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *UNPCKHPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *UNPCKHPS128) Name() string {
	return "UNPCKHPS (128 bit) interleave-high"
}

func (v *UNPCKHPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *UNPCKHPS128) Stub() string {
	return stubUnpckhps128
}

func (v *UNPCKHPS128) Assembly() string {
	return assemblyUnpckhps128
}

func (v *UNPCKHPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	unpckhps128(&floats1, &floats2, &ret)

	log.Printf("UNPCKHPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *UNPCKHPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
