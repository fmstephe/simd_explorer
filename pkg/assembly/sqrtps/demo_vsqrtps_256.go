package sqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsqrtps_256.s
var assemblyVsqrtps256 string

//go:embed stub_vsqrtps_256.go
var stubVsqrtps256 string

type VSQRTPS256 struct {
}

func (v *VSQRTPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VSQRTPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VSQRTPS256) Name() string {
	return "VSQRTPS YMM (256 bit)"
}

func (v *VSQRTPS256) Description() string {
	return "AVX form: compute square root of packed single-precision floats in YMM, lane-wise."
}

func (v *VSQRTPS256) Stub() string {
	return stubVsqrtps256
}

func (v *VSQRTPS256) Assembly() string {
	return assemblyVsqrtps256
}

func (v *VSQRTPS256) Run(inputs [][]byte) (output []byte) {
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vsqrtps256(&floats, &ret)

	log.Printf("VSQRTPS256 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSQRTPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
