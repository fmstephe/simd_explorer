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
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVDQU128() *VMOVDQU128 {
	return &VMOVDQU128{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VMOVDQU128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVDQU128) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVDQU128) Name() string {
	return "VMOVDQU XMM (128 bit)"
}

func (v *VMOVDQU128) Description() string {
	return "Unaligned move of packed 32-bit integers between memory and XMM; copies data unchanged."
}

func (v *VMOVDQU128) Stub() string {
	return stubVmovdqu128
}

func (v *VMOVDQU128) Assembly() string {
	return assemblyVmovdqu128
}

func (v *VMOVDQU128) Run(_ [][]byte) (output []byte) {
	uints := [4]uint32{}
	copy(uints[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [4]uint32{}

	vmovdqu128(&uints, &ret)

	log.Printf("VMOVDQU128 input %v output %v", uints, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVDQU128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
