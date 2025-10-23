package divps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdivps_256.s
var assemblyVdivps256 string

//go:embed stub_vdivps_256.go
var stubVdivps256 string

type VDIVPS256 struct {
}

func (v *VDIVPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VDIVPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VDIVPS256) Name() string {
	return "VDIVPS YMM (256 bit)"
}

func (v *VDIVPS256) Description() string {
	return "TODO"
}

func (v *VDIVPS256) Stub() string {
	return stubVdivps256
}

func (v *VDIVPS256) Assembly() string {
	return assemblyVdivps256
}

func (v *VDIVPS256) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vdivps256(&floats1, &floats2, &ret)

	log.Printf("VDIVPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VDIVPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
