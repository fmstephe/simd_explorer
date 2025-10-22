package rcpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rcpss_128.s
var assemblyRcpss128 string

//go:embed stub_rcpss_128.go
var stubRcpss128 string

type RCPSS128 struct {
}

func (v *RCPSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *RCPSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *RCPSS128) Name() string {
	return "RCPSS (128 bit) "
}

func (v *RCPSS128) Description() string {
	return "TODO"
}

func (v *RCPSS128) Stub() string {
	return stubRcpss128
}

func (v *RCPSS128) Assembly() string {
	return assemblyRcpss128
}

func (v *RCPSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	rcpss128(&floats, &ret)

	log.Printf("RCPSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *RCPSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}