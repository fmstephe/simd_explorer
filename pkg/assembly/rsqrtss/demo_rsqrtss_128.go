package rsqrtss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rsqrtss_128.s
var assemblyRsqrtss128 string

//go:embed stub_rsqrtss_128.go
var stubRsqrtss128 string

type RSQRTSS128 struct {
}

func (v *RSQRTSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *RSQRTSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *RSQRTSS128) Name() string {
	return "RSQRTSS (128 bit) "
}

func (v *RSQRTSS128) Description() string {
	return "Compute reciprocal square root estimate of scalar single-precision (lane 0); upper lanes pass through."
}

func (v *RSQRTSS128) Stub() string {
	return stubRsqrtss128
}

func (v *RSQRTSS128) Assembly() string {
	return assemblyRsqrtss128
}

func (v *RSQRTSS128) Run(inputs [][]byte) (output []byte) {
	// Example arguments processing
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	rsqrtss128(&floats, &ret)

	log.Printf("RSQRTSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *RSQRTSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
