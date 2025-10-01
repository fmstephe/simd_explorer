package addps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_128_addps.s
var assemblyAddps128 string

//go:embed stub_128_addps.go
var stubAddps128 string

type ADDPS128 struct {
}

func (v *ADDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *ADDPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *ADDPS128) Name() string {
	return "ADDPS XMM (128 bit)"
}

func (v *ADDPS128) Description() string {
	return "TODO"
}

func (v *ADDPS128) Stub() string {
	return stubAddps128
}

func (v *ADDPS128) Assembly() string {
	return assemblyAddps128
}

func (v *ADDPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	addps128(&floats1, &floats2, &ret)

	log.Printf("ADDPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *ADDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
