package unpcklps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpcklps_256.s
var assemblyVunpcklps256 string

//go:embed stub_vunpcklps_256.go
var stubVunpcklps256 string

type VUNPCKLPS256 struct {
}

func (v *VUNPCKLPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VUNPCKLPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VUNPCKLPS256) Name() string {
	return "VUNPCKLPS (256 bit) interleave-high"
}

func (v *VUNPCKLPS256) Description() string {
	return "Unpack and interleave high 64-bit halves per 128-bit lane: [a2,b2,a3,b3 | a6,b6,a7,b7]."
}

func (v *VUNPCKLPS256) Stub() string {
	return stubVunpcklps256
}

func (v *VUNPCKLPS256) Assembly() string {
	return assemblyVunpcklps256
}

func (v *VUNPCKLPS256) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vunpcklps256(&floats1, &floats2, &ret)

	log.Printf("VUNPCKLPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VUNPCKLPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
