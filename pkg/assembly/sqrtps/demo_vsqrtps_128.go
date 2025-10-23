package sqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsqrtps_128.s
var assemblyVsqrtps128 string

//go:embed stub_vsqrtps_128.go
var stubVsqrtps128 string

type VSQRTPS128 struct {
}

func (v *VSQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VSQRTPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VSQRTPS128) Name() string {
	return "VSQRTPS XMM (128 bit)"
}

func (v *VSQRTPS128) Description() string {
	return "TODO"
}

func (v *VSQRTPS128) Stub() string {
	return stubVsqrtps128
}

func (v *VSQRTPS128) Assembly() string {
	return assemblyVsqrtps128
}

func (v *VSQRTPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vsqrtps128(&floats, &ret)

	log.Printf("VSQRTPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VSQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
