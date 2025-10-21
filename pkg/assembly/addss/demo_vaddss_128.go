package addss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddss_128.s
var assemblyVaddss128 string

//go:embed stub_vaddss_128.go
var stubVaddss128 string

type VADDSS128 struct {
}

func (v *VADDSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VADDSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VADDSS128) Name() string {
	return "VADDSS (128 bit) "
}

func (v *VADDSS128) Description() string {
	return "TODO"
}

func (v *VADDSS128) Stub() string {
	return stubVaddss128
}

func (v *VADDSS128) Assembly() string {
	return assemblyVaddss128
}

func (v *VADDSS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	addss128(&floats1, &floats2, &ret)

	log.Printf("VADDSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VADDSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
