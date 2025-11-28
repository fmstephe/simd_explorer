package divss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdivss_128.s
var assemblyVdivss128 string

//go:embed stub_vdivss_128.go
var stubVdivss128 string

type VDIVSS128 struct {
}

func (v *VDIVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VDIVSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VDIVSS128) Name() string {
	return "VDIVSS (128 bit) "
}

func (v *VDIVSS128) Description() string {
	return "AVX form: divide scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VDIVSS128) Stub() string {
	return stubVdivss128
}

func (v *VDIVSS128) Assembly() string {
	return assemblyVdivss128
}

func (v *VDIVSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vdivss128(&floats1, &floats2, &ret)

	log.Printf("VDIVSS128 input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VDIVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
