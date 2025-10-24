package rsqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrsqrtps_128.s
var assemblyVrsqrtps128 string

//go:embed stub_vrsqrtps_128.go
var stubVrsqrtps128 string

type VRSQRTPS128 struct {
}

func (v *VRSQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VRSQRTPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VRSQRTPS128) Name() string {
	return "VRSQRTPS XMM (128 bit)"
}

func (v *VRSQRTPS128) Description() string {
	return "TODO"
}

func (v *VRSQRTPS128) Stub() string {
	return stubVrsqrtps128
}

func (v *VRSQRTPS128) Assembly() string {
	return assemblyVrsqrtps128
}

func (v *VRSQRTPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vrsqrtps128(&floats, &ret)

	log.Printf("VRSQRTPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VRSQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
