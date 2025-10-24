package minps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vminps_128.s
var assemblyVminps128 string

//go:embed stub_vminps_128.go
var stubVminps128 string

type VMINPS128 struct {
}

func (v *VMINPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMINPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMINPS128) Name() string {
	return "VMINPS XMM (128 bit)"
}

func (v *VMINPS128) Description() string {
	return "TODO"
}

func (v *VMINPS128) Stub() string {
	return stubVminps128
}

func (v *VMINPS128) Assembly() string {
	return assemblyVminps128
}

func (v *VMINPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vminps128(&floats1, &floats2, &ret)

	log.Printf("VMINPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMINPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
