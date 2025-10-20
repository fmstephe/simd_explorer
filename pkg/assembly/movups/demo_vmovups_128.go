package movups

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovups_128.s
var assemblyVmovups128 string

//go:embed stub_vmovups_128.go
var stubVmovups128 string

type VMOVUPS128 struct {
}

func (v *VMOVUPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VMOVUPS128) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
}

func (v *VMOVUPS128) Name() string {
	return "VMOVUPS XMM (128 bit)"
}

func (v *VMOVUPS128) Description() string {
	return "TODO"
}

func (v *VMOVUPS128) Stub() string {
	return stubVmovups128
}

func (v *VMOVUPS128) Assembly() string {
	return assemblyVmovups128
}

func (v *VMOVUPS128) Run(inputs [][]byte) (output []byte) {
	floats := [4]float32{}
	copy(floats[:], number.ToFloat32Slice(inputs[0]))

	ret := [4]float32{}

	vmovups128(&floats, &ret)

	log.Printf("VMOVUPS128 input %v output %v", floats, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VMOVUPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
