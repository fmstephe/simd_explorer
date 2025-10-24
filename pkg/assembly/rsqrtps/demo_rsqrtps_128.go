package rsqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rsqrtps_128.s
var assemblyRsqrtps128 string

//go:embed stub_rsqrtps_128.go
var stubRsqrtps128 string

type RSQRTPS128 struct {
}

func (v *RSQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *RSQRTPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *RSQRTPS128) Name() string {
	return "RSQRTPS XMM (128 bit)"
}

func (v *RSQRTPS128) Description() string {
	return "TODO"
}

func (v *RSQRTPS128) Stub() string {
	return stubRsqrtps128
}

func (v *RSQRTPS128) Assembly() string {
	return assemblyRsqrtps128
}

func (v *RSQRTPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	rsqrtps128(&floats, &ret)

	log.Printf("RSQRTPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *RSQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
