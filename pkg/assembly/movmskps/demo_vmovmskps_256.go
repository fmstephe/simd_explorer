package movmskps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskps_256.s
var assemblyVmovmskps256 string

//go:embed stub_vmovmskps_256.go
var stubVmovmskps256 string

type VMOVMSKPS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVMSKPS256() *VMOVMSKPS256 {
	return &VMOVMSKPS256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VMOVMSKPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVMSKPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVMSKPS256) Name() string {
	return "VMOVMSKPS (256 bit) "
}

func (v *VMOVMSKPS256) Description() string {
	return "Extract sign bits of packed single-precision elements in YMM into an 8-bit integer mask."
}

func (v *VMOVMSKPS256) Stub() string {
	return stubVmovmskps256
}

func (v *VMOVMSKPS256) Assembly() string {
	return assemblyVmovmskps256
}

func (v *VMOVMSKPS256) Run() {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]byte{}

	vmovmskps256(&vals, &ret)

	log.Printf("VMOVMSKPS256 input %v output %v", vals, ret)

	out := ret[:]
	v.ret.SetData(out)

}

func (v *VMOVMSKPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
