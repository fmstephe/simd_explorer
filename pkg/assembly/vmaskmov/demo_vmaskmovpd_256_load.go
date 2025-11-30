package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovpd_256_load.s
var assemblyVmaskmovpd256Load string

//go:embed stub_vmaskmovpd_256_load.go
var stubVmaskmovpd256Load string

type VMASKMOVPD256LOAD struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPD256LOAD() *VMASKMOVPD256LOAD {
	return &VMASKMOVPD256LOAD{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		mask: number.NewNamedUintParameter("mask", 256, 64, 16),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VMASKMOVPD256LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPD256LOAD) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPD256LOAD) Name() string {
	return "VMASKMOVPD (256 bit) load"
}

func (v *VMASKMOVPD256LOAD) Description() string {
	return "Masked load of packed double-precision floats (per 128-bit lane); unselected lanes are zeroed."
}

func (v *VMASKMOVPD256LOAD) Stub() string {
	return stubVmaskmovpd256Load
}

func (v *VMASKMOVPD256LOAD) Assembly() string {
	return assemblyVmaskmovpd256Load
}

func (v *VMASKMOVPD256LOAD) Run() {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	mask := [4]float64{}
	copy(mask[:], number.ToFloat64Slice(v.mask.FlatData()))

	ret := [4]float64{}

	vmaskmovpd256Load(&vals, &mask, &ret)

	log.Printf("VMASKMOVPD256LOAD vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VMASKMOVPD256LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
