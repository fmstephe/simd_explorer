package unpcklps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpcklps_128.s
var assemblyVunpcklps128 string

//go:embed stub_vunpcklps_128.go
var stubVunpcklps128 string

type VUNPCKLPS128 struct {
}

func (v *VUNPCKLPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VUNPCKLPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VUNPCKLPS128) Name() string {
	return "VUNPCKLPS (128 bit) interleave-high"
}

func (v *VUNPCKLPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *VUNPCKLPS128) Stub() string {
	return stubVunpcklps128
}

func (v *VUNPCKLPS128) Assembly() string {
	return assemblyVunpcklps128
}

func (v *VUNPCKLPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vunpcklps128(&floats1, &floats2, &ret)

	log.Printf("VUNPCKLPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VUNPCKLPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
