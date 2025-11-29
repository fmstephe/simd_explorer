package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_128_store.s
var assemblyVmaskmovpd128Store string

//go:embed stub_vmaskmovpd_128_store.go
var stubVmaskmovpd128Store string

type VMASKMOVPD128STORE struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPD128STORE() *VMASKMOVPD128STORE {
	return &VMASKMOVPD128STORE{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		mask: number.NewNamedUintParameter("mask", 128, 64, 16),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VMASKMOVPD128STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPD128STORE) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPD128STORE) Name() string {
	return "VMASKMOVPD (128 bit) store"
}

func (v *VMASKMOVPD128STORE) Description() string {
	return "Masked store of packed double-precision floats: store lanes where mask sign-bit is set."
}

func (v *VMASKMOVPD128STORE) Stub() string {
	return stubVmaskmovpd128Store
}

func (v *VMASKMOVPD128STORE) Assembly() string {
	return assemblyVmaskmovpd128Store
}

func (v *VMASKMOVPD128STORE) Run(_ [][]byte) (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [2]float64{}

	vmaskmovpd128Store(&vals, &mask, &ret)

	log.Printf("VMASKMOVPD128STORE vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMASKMOVPD128STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
