package vmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vmaskmovps_128_store.s
var assemblyVmaskmovps128Store string

//go:embed stub_vmaskmovps_128_store.go
var stubVmaskmovps128Store string

type VMASKMOVPS128STORE struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVMASKMOVPS128STORE() *VMASKMOVPS128STORE {
	return &VMASKMOVPS128STORE{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		mask: number.NewNamedUintParameter("mask", 128, 32, 16),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMASKMOVPS128STORE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VMASKMOVPS128STORE) Output() *number.Parameter {
	return v.ret
}

func (v *VMASKMOVPS128STORE) Name() string {
	return "VMASKMOVPS (128 bit) store"
}

func (v *VMASKMOVPS128STORE) Description() string {
	return "Masked store of packed single-precision floats: store lanes where mask sign-bit is set."
}

func (v *VMASKMOVPS128STORE) Stub() string {
	return stubVmaskmovps128Store
}

func (v *VMASKMOVPS128STORE) Assembly() string {
	return assemblyVmaskmovps128Store
}

func (v *VMASKMOVPS128STORE) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	mask := [4]float32{}
	copy(mask[:], number.ToFloat32Slice(v.mask.FlatData()))

	ret := [4]float32{}

	vmaskmovps128Store(&vals, &mask, &ret)

	log.Printf("VMASKMOVPS128STORE vals %v mask %v ret %v", vals, mask, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMASKMOVPS128STORE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
