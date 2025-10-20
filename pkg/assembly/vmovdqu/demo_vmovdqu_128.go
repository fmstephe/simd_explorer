package vmovdqu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovdqu_128.s
var assemblyVmovdqu128 string

//go:embed stub_vmovdqu_128.go
var stubVmovdqu128 string

type VMOVDQU128 struct {
}

func (v *VMOVDQU128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewUintParameter(128, 32, 10),
	}
}

func (v *VMOVDQU128) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 10)
}

func (v *VMOVDQU128) Name() string {
	return "VMOVDQU XMM (128 bit)"
}

func (v *VMOVDQU128) Description() string {
	return "TODO"
}

func (v *VMOVDQU128) Stub() string {
	return stubVmovdqu128
}

func (v *VMOVDQU128) Assembly() string {
	return assemblyVmovdqu128
}

func (v *VMOVDQU128) Run(inputs [][]byte) (output []byte) {
	uints := [4]uint32{}
	copy(uints[:], number.ToUint32Slice(inputs[0]))

	ret := [4]uint32{}

	vmovdqu128(&uints, &ret)

	log.Printf("VMOVDQU128 input %v output %v", uints, ret)

	return number.Uint32SliceToBytes(ret[:])
}

func (v *VMOVDQU128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
