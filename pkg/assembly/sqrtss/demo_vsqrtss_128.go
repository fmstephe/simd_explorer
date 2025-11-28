package sqrtss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsqrtss_128.s
var assemblyVsqrtss128 string

//go:embed stub_vsqrtss_128.go
var stubVsqrtss128 string

type VSQRTSS128 struct {
}

func (v *VSQRTSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSQRTSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSQRTSS128) Name() string {
	return "VSQRTSS (128 bit) "
}

func (v *VSQRTSS128) Description() string {
	return "AVX form: compute square root of scalar single-precision (lane 0); upper lanes pass through."
}

func (v *VSQRTSS128) Stub() string {
	return stubVsqrtss128
}

func (v *VSQRTSS128) Assembly() string {
	return assemblyVsqrtss128
}

func (v *VSQRTSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vsqrtss128(&floats, &ret)

	log.Printf("VSQRTSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSQRTSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
