package unpcklps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_unpcklps_128.s
var assemblyUnpcklps128 string

//go:embed stub_unpcklps_128.go
var stubUnpcklps128 string

type UNPCKLPS128 struct {
}

func (v *UNPCKLPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *UNPCKLPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *UNPCKLPS128) Name() string {
	return "UNPCKLPS (128 bit) interleave-high"
}

func (v *UNPCKLPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *UNPCKLPS128) Stub() string {
	return stubUnpcklps128
}

func (v *UNPCKLPS128) Assembly() string {
	return assemblyUnpcklps128
}

func (v *UNPCKLPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	unpcklps128(&floats1, &floats2, &ret)

	log.Printf("UNPCKLPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *UNPCKLPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
