package minps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_minps_128.s
var assemblyMinps128 string

//go:embed stub_minps_128.go
var stubMinps128 string

type MINPS128 struct {
}

func (v *MINPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *MINPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *MINPS128) Name() string {
	return "MINPS XMM (128 bit)"
}

func (v *MINPS128) Description() string {
	return "TODO"
}

func (v *MINPS128) Stub() string {
	return stubMinps128
}

func (v *MINPS128) Assembly() string {
	return assemblyMinps128
}

func (v *MINPS128) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	minps128(&floats1, &floats2, &ret)

	log.Printf("MINPS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *MINPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
