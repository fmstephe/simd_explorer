package minss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vminss_128.s
var assemblyVminss128 string

//go:embed stub_vminss_128.go
var stubVminss128 string

type VMINSS128 struct {
}

func (v *VMINSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMINSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMINSS128) Name() string {
	return "VMINSS (128 bit) "
}

func (v *VMINSS128) Description() string {
	return "AVX form: compute minimum of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VMINSS128) Stub() string {
	return stubVminss128
}

func (v *VMINSS128) Assembly() string {
	return assemblyVminss128
}

func (v *VMINSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vminss128(&floats1, &floats2, &ret)

	log.Printf("VMINSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMINSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
