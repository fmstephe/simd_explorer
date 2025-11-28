package vmovdqa

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovdqa_128.s
var assemblyVmovdqa128 string

//go:embed stub_vmovdqa_128.go
var stubVmovdqa128 string

type VMOVDQA128 struct {
}

func (v *VMOVDQA128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 32, 10),
	}
}

func (v *VMOVDQA128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10)
}

func (v *VMOVDQA128) Name() string {
	return "VMOVDQA XMM (128 bit)"
}

func (v *VMOVDQA128) Description() string {
	return "Aligned move of packed 32-bit integers between memory and XMM; copies data unchanged."
}

func (v *VMOVDQA128) Stub() string {
	return stubVmovdqa128
}

func (v *VMOVDQA128) Assembly() string {
	return assemblyVmovdqa128
}

func (v *VMOVDQA128) Run(inputs [][]byte) (output []byte) {
	uints := [4]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [4]uint32{}

	vmovdqa128(&uints, &ret)

	log.Printf("VMOVDQA128 input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQA128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
