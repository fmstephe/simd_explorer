package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_128_load.s
var assemblyVmaskmovps128Load string

//go:embed stub_vmaskmovps_128_load.go
var stubVmaskmovps128Load string

type VMASKMOVPS128LOAD struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPS128LOAD() *VMASKMOVPS128LOAD {
	return &VMASKMOVPS128LOAD{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		mask: number.NewNamedUintParameter("mask", 128, 32, 16),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMASKMOVPS128LOAD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPS128LOAD) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPS128LOAD) Name() string {
	return "VMASKMOVPS (128 bit) load"
}

func (v *VMASKMOVPS128LOAD) Description() string {
	return "Masked load of packed single-precision floats; unselected lanes are zeroed."
}

func (v *VMASKMOVPS128LOAD) Stub() string {
	return stubVmaskmovps128Load
}

func (v *VMASKMOVPS128LOAD) Assembly() string {
	return assemblyVmaskmovps128Load
}

func (v *VMASKMOVPS128LOAD) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	mask := [4]float32{}
	copy(mask[:], number.ToFloat32Slice(v.mask.FlatData()))

	ret := [4]float32{}

	vmaskmovps128Load(&vals, &mask, &ret)

	log.Printf("VMASKMOVPS128LOAD vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMASKMOVPS128LOAD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
