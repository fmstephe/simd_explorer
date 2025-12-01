package vpermd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermd_256.s
var assemblyVpermd256 string

//go:embed stub_vpermd_256.go
var stubVpermd256 string

type VPERMD256 struct {
	vals    *number.Parameter
	control *number.Parameter
	ret     *number.Parameter
}

func NewVPERMD256() *VPERMD256 {
	return &VPERMD256{
		vals:    number.NewNamedUintParameter("vals", 256, 32, 16),
		control: number.NewNamedUintParameter("control", 256, 32, 10),
		ret:     number.NewNamedUintParameter("ret", 256, 32, 16),
	}
}

func (v *VPERMD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.control,
	}
}

func (v *VPERMD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMD256) Name() string {
	return "VPERMD (256 bit) "
}

func (v *VPERMD256) Description() string {
	return "Permute 8 u32 elements from vals using per-dword indices in control (AVX2 VPERMD; indices select within 128-bit lanes)."
}

func (v *VPERMD256) Stub() string {
	return stubVpermd256
}

func (v *VPERMD256) Assembly() string {
	return assemblyVpermd256
}

func (v *VPERMD256) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	control := [8]uint32{}
	copy(control[:], number.ToUint32Slice(v.control.FlatData()))

	ret := [8]uint32{}

	vpermd256(&vals, &control, &ret)

	log.Printf("VPERMD256 control %v vals %v ret %v", control, vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
