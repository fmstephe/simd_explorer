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
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVDQA128() *VMOVDQA128 {
	return &VMOVDQA128{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VMOVDQA128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVDQA128) Output() *number.Parameter {
	return v.ret
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

func (v *VMOVDQA128) Run() {
	uints := [4]uint32{}
	copy(uints[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [4]uint32{}

	vmovdqa128(&uints, &ret)

	log.Printf("VMOVDQA128 input %v output %v", uints, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VMOVDQA128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
