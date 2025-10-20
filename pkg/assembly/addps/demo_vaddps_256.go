package addps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddps_256.s
var assemblyVaddps256 string

//go:embed stub_vaddps_256.go
var stubVaddps256 string

type VADDPS256 struct {
}

func (v *VADDPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VADDPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VADDPS256) Name() string {
	return "VADDPS YMM (256 bit)"
}

func (v *VADDPS256) Description() string {
	return "TODO"
}

func (v *VADDPS256) Stub() string {
	return stubVaddps256
}

func (v *VADDPS256) Assembly() string {
	return assemblyVaddps256
}

func (v *VADDPS256) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vaddps256(&floats1, &floats2, &ret)

	log.Printf("VADDPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VADDPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
