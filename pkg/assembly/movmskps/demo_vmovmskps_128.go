package movmskps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskps_128.s
var assemblyVmovmskps128 string

//go:embed stub_vmovmskps_128.go
var stubVmovmskps128 string

type VMOVMSKPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVMSKPS128() *VMOVMSKPS128 {
	return &VMOVMSKPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VMOVMSKPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVMSKPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVMSKPS128) Name() string {
	return "VMOVMSKPS (128 bit) "
}

func (v *VMOVMSKPS128) Description() string {
	return "Extract sign bits of packed single-precision elements in XMM into a 4-bit integer mask."
}

func (v *VMOVMSKPS128) Stub() string {
	return stubVmovmskps128
}

func (v *VMOVMSKPS128) Assembly() string {
	return assemblyVmovmskps128
}

func (v *VMOVMSKPS128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]byte{}

	vmovmskps128(&vals, &ret)

	log.Printf("VMOVMSKPS128 input %v output %v", vals, ret)

	out := ret[:]
	v.ret.SetData(out)

}

func (v *VMOVMSKPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
