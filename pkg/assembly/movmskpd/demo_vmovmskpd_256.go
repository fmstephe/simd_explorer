package movmskpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskpd_256.s
var assemblyVmovmskpd256 string

//go:embed stub_vmovmskpd_256.go
var stubVmovmskpd256 string

type VMOVMSKPD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVMSKPD256() *VMOVMSKPD256 {
	return &VMOVMSKPD256{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VMOVMSKPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVMSKPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVMSKPD256) Name() string {
	return "VMOVMSKPD (256 bit) "
}

func (v *VMOVMSKPD256) Description() string {
	return "Extract sign bits of packed double-precision elements in YMM into a 4-bit integer mask."
}

func (v *VMOVMSKPD256) Stub() string {
	return stubVmovmskpd256
}

func (v *VMOVMSKPD256) Assembly() string {
	return assemblyVmovmskpd256
}

func (v *VMOVMSKPD256) Run() {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]byte{}

	vmovmskpd256(&vals, &ret)

	log.Printf("VMOVMSKPD256 input %v output %v", vals, ret)

	out := ret[:]
	v.ret.SetData(out)

}

func (v *VMOVMSKPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
