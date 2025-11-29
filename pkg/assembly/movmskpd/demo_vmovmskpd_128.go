package movmskpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmovmskpd_128.s
var assemblyVmovmskpd128 string

//go:embed stub_vmovmskpd_128.go
var stubVmovmskpd128 string

type VMOVMSKPD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVMOVMSKPD128() *VMOVMSKPD128 {
	return &VMOVMSKPD128{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *VMOVMSKPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VMOVMSKPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VMOVMSKPD128) Name() string {
	return "VMOVMSKPD (128 bit) "
}

func (v *VMOVMSKPD128) Description() string {
	return "Extract sign bits of packed double-precision elements in XMM into a 2-bit integer mask."
}

func (v *VMOVMSKPD128) Stub() string {
	return stubVmovmskpd128
}

func (v *VMOVMSKPD128) Assembly() string {
	return assemblyVmovmskpd128
}

func (v *VMOVMSKPD128) Run(_ [][]byte) (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]byte{}

	vmovmskpd128(&vals, &ret)

	log.Printf("VMOVMSKPD128 input %v output %v", vals, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *VMOVMSKPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
