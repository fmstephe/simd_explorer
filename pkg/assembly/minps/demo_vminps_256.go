package minps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vminps_256.s
var assemblyVminps256 string

//go:embed stub_vminps_256.go
var stubVminps256 string

type VMINPS256 struct {
}

func (v *VMINPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMINPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMINPS256) Name() string {
	return "VMINPS YMM (256 bit)"
}

func (v *VMINPS256) Description() string {
	return "TODO"
}

func (v *VMINPS256) Stub() string {
	return stubVminps256
}

func (v *VMINPS256) Assembly() string {
	return assemblyVminps256
}

func (v *VMINPS256) Run(inputs [][]byte) (output []byte) {
	floats1 := [8]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [8]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [8]float32{}

	vminps256(&floats1, &floats2, &ret)

	log.Printf("VMINPS256 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMINPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
