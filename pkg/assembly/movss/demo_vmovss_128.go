package movss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovss_128.s
var assemblyVmovss128 string

//go:embed stub_vmovss_128.go
var stubVmovss128 string

type VMOVSS128 struct {
}

func (v *VMOVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVSS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVSS128) Name() string {
	return "VMOVSS XMM (128 bit)"
}

func (v *VMOVSS128) Description() string {
	return "AVX form: move scalar single-precision (lane 0) between XMM and memory; upper lanes pass through."
}

func (v *VMOVSS128) Stub() string {
	return stubVmovss128
}

func (v *VMOVSS128) Assembly() string {
	return assemblyVmovss128
}

func (v *VMOVSS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vmovss128(&floats, &ret)

	log.Printf("VMOVSS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
