package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_128_load.s
var assemblyVmaskmovpd128Load string

//go:embed stub_vmaskmovpd_128_load.go
var stubVmaskmovpd128Load string

type VMASKMOVPD128LOAD struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPD128LOAD() *VMASKMOVPD128LOAD {
	return &VMASKMOVPD128LOAD{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		mask: number.NewNamedUintParameter("mask", 128, 64, 16),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VMASKMOVPD128LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPD128LOAD) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPD128LOAD) Name() string {
	return "VMASKMOVPD (128 bit) load"
}

func (v *VMASKMOVPD128LOAD) Description() string {
	return "Masked load of packed double-precision floats; unselected lanes are zeroed."
}

func (v *VMASKMOVPD128LOAD) Stub() string {
	return stubVmaskmovpd128Load
}

func (v *VMASKMOVPD128LOAD) Assembly() string {
	return assemblyVmaskmovpd128Load
}

func (v *VMASKMOVPD128LOAD) Run() (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	mask := [2]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [2]float64{}

	vmaskmovpd128Load(&vals, &mask, &ret)

	log.Printf("VMASKMOVPD128LOAD vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMASKMOVPD128LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
