package maxps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_maxps_128.s
var assemblyMaxps128 string

//go:embed stub_maxps_128.go
var stubMaxps128 string

type MAXPS128 struct {
}

func (v *MAXPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *MAXPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MAXPS128) Name() string {
	return "MAXPS XMM (128 bit)"
}

func (v *MAXPS128) Description() string {
	return "TODO"
}

func (v *MAXPS128) Stub() string {
	return stubMaxps128
}

func (v *MAXPS128) Assembly() string {
	return assemblyMaxps128
}

func (v *MAXPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	maxps128(&floats1, &floats2, &ret)

	log.Printf("MAXPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MAXPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
