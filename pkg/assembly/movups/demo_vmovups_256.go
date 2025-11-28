package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovups_256.s
var assemblyVmovups256 string

//go:embed stub_vmovups_256.go
var stubVmovups256 string

type VMOVUPS256 struct {
}

func (v *VMOVUPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(256, 32),
	}
}

func (v *VMOVUPS256) Output() *number.Parameter {
	return number.NewFloatParameter(256, 32)
}

func (v *VMOVUPS256) Name() string {
	return "VMOVUPS YMM (256 bit)"
}

func (v *VMOVUPS256) Description() string {
	return "Unaligned move of packed single-precision floats between memory and YMM; copies data unchanged."
}

func (v *VMOVUPS256) Stub() string {
	return stubVmovups256
}

func (v *VMOVUPS256) Assembly() string {
	return assemblyVmovups256
}

func (v *VMOVUPS256) Run(inputs [][]byte) (output []byte) {
	floats := [8]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [8]float32{}

	vmovups256(&floats, &ret)

	log.Printf("VMOVUPS256 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVUPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
