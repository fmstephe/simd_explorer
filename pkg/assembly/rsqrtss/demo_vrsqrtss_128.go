package rsqrtss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrsqrtss_128.s
var assemblyVrsqrtss128 string

//go:embed stub_vrsqrtss_128.go
var stubVrsqrtss128 string

type VRSQRTSS128 struct {
}

func (v *VRSQRTSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VRSQRTSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VRSQRTSS128) Name() string {
	return "VRSQRTSS (128 bit) "
}

func (v *VRSQRTSS128) Description() string {
	return "AVX form: reciprocal square root estimate of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VRSQRTSS128) Stub() string {
	return stubVrsqrtss128
}

func (v *VRSQRTSS128) Assembly() string {
	return assemblyVrsqrtss128
}

func (v *VRSQRTSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vrsqrtss128(&floats, &ret)

	log.Printf("VRSQRTSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VRSQRTSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
