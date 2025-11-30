package vmovdqu

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovdqu_256.s
var assemblyVmovdqu256 string

//go:embed stub_vmovdqu_256.go
var stubVmovdqu256 string

type VMOVDQU256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVDQU256() *VMOVDQU256 {
	return &VMOVDQU256{
		vals: number.NewNamedUintParameter("vals", 256, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VMOVDQU256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVDQU256) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVDQU256) Name() string {
	return "VMOVDQU YMM (256 bit)"
}

func (v *VMOVDQU256) Description() string {
	return "Unaligned move of packed 32-bit integers between memory and YMM; copies data unchanged."
}

func (v *VMOVDQU256) Stub() string {
	return stubVmovdqu256
}

func (v *VMOVDQU256) Assembly() string {
	return assemblyVmovdqu256
}

func (v *VMOVDQU256) Run() (output []byte) {
	uints := [8]uint32{}
	copy(uints[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [8]uint32{}

	vmovdqu256(&uints, &ret)

	log.Printf("VMOVDQU256 input %v output %v", uints, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMOVDQU256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
