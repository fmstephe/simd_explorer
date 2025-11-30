package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_256_load.s
var assemblyVmaskmovps256Load string

//go:embed stub_vmaskmovps_256_load.go
var stubVmaskmovps256Load string

type VMASKMOVPS256LOAD struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPS256LOAD() *VMASKMOVPS256LOAD {
	return &VMASKMOVPS256LOAD{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		mask: number.NewNamedUintParameter("mask", 256, 32, 16),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMASKMOVPS256LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPS256LOAD) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPS256LOAD) Name() string {
	return "VMASKMOVPS (256 bit) load"
}

func (v *VMASKMOVPS256LOAD) Description() string {
	return "Masked load of packed single-precision floats (per 128-bit lane); unselected lanes are zeroed."
}

func (v *VMASKMOVPS256LOAD) Stub() string {
	return stubVmaskmovps256Load
}

func (v *VMASKMOVPS256LOAD) Assembly() string {
	return assemblyVmaskmovps256Load
}

func (v *VMASKMOVPS256LOAD) Run() (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	mask := [8]float32{}
	copy(mask[:], number.ToFloat32Slice(v.mask.FlatData()))

	ret := [8]float32{}

	vmaskmovps256Load(&vals, &mask, &ret)

	log.Printf("VMASKMOVPS256LOAD vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMASKMOVPS256LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
