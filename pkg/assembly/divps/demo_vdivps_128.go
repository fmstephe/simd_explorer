package divps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdivps_128.s
var assemblyVdivps128 string

//go:embed stub_vdivps_128.go
var stubVdivps128 string

type VDIVPS128 struct {
}

func (v *VDIVPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VDIVPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VDIVPS128) Name() string {
	return "VDIVPS XMM (128 bit)"
}

func (v *VDIVPS128) Description() string {
	return "TODO"
}

func (v *VDIVPS128) Stub() string {
	return stubVdivps128
}

func (v *VDIVPS128) Assembly() string {
	return assemblyVdivps128
}

func (v *VDIVPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vdivps128(&floats1, &floats2, &ret)

	log.Printf("VDIVPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VDIVPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
