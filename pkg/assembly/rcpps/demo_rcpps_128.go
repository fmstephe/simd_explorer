package rcpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rcpps_128.s
var assemblyRcpps128 string

//go:embed stub_rcpps_128.go
var stubRcpps128 string

type RCPPS128 struct {
}

func (v *RCPPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *RCPPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *RCPPS128) Name() string {
	return "RCPPS XMM (128 bit)"
}

func (v *RCPPS128) Description() string {
	return "TODO"
}

func (v *RCPPS128) Stub() string {
	return stubRcpps128
}

func (v *RCPPS128) Assembly() string {
	return assemblyRcpps128
}

func (v *RCPPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	rcpps128(&floats, &ret)

	log.Printf("RCPPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *RCPPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
