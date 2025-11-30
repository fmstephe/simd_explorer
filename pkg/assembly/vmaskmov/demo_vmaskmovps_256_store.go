package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_256_store.s
var assemblyVmaskmovps256Store string

//go:embed stub_vmaskmovps_256_store.go
var stubVmaskmovps256Store string

type VMASKMOVPS256STORE struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPS256STORE() *VMASKMOVPS256STORE {
	return &VMASKMOVPS256STORE{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		mask: number.NewNamedUintParameter("mask", 256, 32, 16),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMASKMOVPS256STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPS256STORE) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPS256STORE) Name() string {
	return "VMASKMOVPS (256 bit) store"
}

func (v *VMASKMOVPS256STORE) Description() string {
	return "Masked store of packed single-precision floats (per 128-bit lane)."
}

func (v *VMASKMOVPS256STORE) Stub() string {
	return stubVmaskmovps256Store
}

func (v *VMASKMOVPS256STORE) Assembly() string {
	return assemblyVmaskmovps256Store
}

func (v *VMASKMOVPS256STORE) Run() {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	mask := [8]float32{}
	copy(mask[:], number.ToFloat32Slice(v.mask.FlatData()))

	ret := [8]float32{}

	vmaskmovps256Store(&vals, &mask, &ret)

	log.Printf("VMASKMOVPS256STORE vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VMASKMOVPS256STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
