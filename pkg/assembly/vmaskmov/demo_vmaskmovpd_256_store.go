package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_256_store.s
var assemblyVmaskmovpd256Store string

//go:embed stub_vmaskmovpd_256_store.go
var stubVmaskmovpd256Store string

type VMASKMOVPD256STORE struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPD256STORE() *VMASKMOVPD256STORE {
	return &VMASKMOVPD256STORE{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		mask: number.NewNamedUintParameter("mask", 256, 64, 16),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VMASKMOVPD256STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPD256STORE) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPD256STORE) Name() string {
	return "VMASKMOVPD (256 bit) store"
}

func (v *VMASKMOVPD256STORE) Description() string {
	return "Masked store of packed double-precision floats (per 128-bit lane)."
}

func (v *VMASKMOVPD256STORE) Stub() string {
	return stubVmaskmovpd256Store
}

func (v *VMASKMOVPD256STORE) Assembly() string {
	return assemblyVmaskmovpd256Store
}

func (v *VMASKMOVPD256STORE) Run(_ [][]byte) (output []byte) {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [4]float64{}

	vmaskmovpd256Store(&vals, &mask, &ret)

	log.Printf("VMASKMOVPD256STORE vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMASKMOVPD256STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
