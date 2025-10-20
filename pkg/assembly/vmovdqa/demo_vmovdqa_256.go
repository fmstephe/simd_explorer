package vmovdqa

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovdqa_256.s
var assemblyVmovdqa256 string

//go:embed stub_vmovdqa_256.go
var stubVmovdqa256 string

type VMOVDQA256 struct {
}

func (v *VMOVDQA256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(256, 32, 10),
	}
}

func (v *VMOVDQA256) Output() *number.Parameter {
	return number.NewUintParameter(256, 32, 10)
}

func (v *VMOVDQA256) Name() string {
	return "VMOVDQA YMM (256 bit)"
}

func (v *VMOVDQA256) Description() string {
	return "TODO"
}

func (v *VMOVDQA256) Stub() string {
	return stubVmovdqa256
}

func (v *VMOVDQA256) Assembly() string {
	return assemblyVmovdqa256
}

func (v *VMOVDQA256) Run(inputs [][]byte) (output []byte) {
	uints := [8]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [8]uint32{}

	vmovdqa256(&uints, &ret)

	log.Printf("VMOVDQA256 input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQA256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
