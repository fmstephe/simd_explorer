package addps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddps_128.s
var assemblyVaddps128 string

//go:embed stub_vaddps_128.go
var stubVaddps128 string

type VADDPS128 struct {
}

func (v *VADDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VADDPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VADDPS128) Name() string {
	return "VADDPS XMM (128 bit)"
}

func (v *VADDPS128) Description() string {
	return "TODO"
}

func (v *VADDPS128) Stub() string {
	return stubVaddps128
}

func (v *VADDPS128) Assembly() string {
	return assemblyVaddps128
}

func (v *VADDPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vaddps128(&floats1, &floats2, &ret)

	log.Printf("VADDPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VADDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
