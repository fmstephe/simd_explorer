package movmskpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_movmskpd_128.s
var assemblyMovmskpd128 string

//go:embed stub_movmskpd_128.go
var stubMovmskpd128 string

type MOVMSKPD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewMOVMSKPD128() *MOVMSKPD128 {
	return &MOVMSKPD128{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedUintParameter("ret", 32, 32, 16),
	}
}

func (v *MOVMSKPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *MOVMSKPD128) Output() *number.Parameter {
	return v.ret
}

func (v *MOVMSKPD128) Name() string {
	return "MOVMSKPD (128 bit) "
}

func (v *MOVMSKPD128) Description() string {
	return "Extract sign bits of packed double-precision elements in XMM into a 2-bit integer mask."
}

func (v *MOVMSKPD128) Stub() string {
	return stubMovmskpd128
}

func (v *MOVMSKPD128) Assembly() string {
	return assemblyMovmskpd128
}

func (v *MOVMSKPD128) Run() (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]byte{}

	movmskpd128(&vals, &ret)

	log.Printf("MOVMSKPD128 input %v output %v", vals, ret)

	out := ret[:]
	v.ret.SetData(out)
	return out
}

func (v *MOVMSKPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
