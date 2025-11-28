package movaps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovaps_128.s
var assemblyVmovaps128 string

//go:embed stub_vmovaps_128.go
var stubVmovaps128 string

type VMOVAPS128 struct {
}

func (v *VMOVAPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVAPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVAPS128) Name() string {
	return "VMOVAPS XMM (128 bit)"
}

func (v *VMOVAPS128) Description() string {
	return "Aligned move of packed single-precision floats between memory and XMM; copies data unchanged."
}

func (v *VMOVAPS128) Stub() string {
	return stubVmovaps128
}

func (v *VMOVAPS128) Assembly() string {
	return assemblyVmovaps128
}

func (v *VMOVAPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vmovaps128(&floats, &ret)

	log.Printf("VMOVAPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVAPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
