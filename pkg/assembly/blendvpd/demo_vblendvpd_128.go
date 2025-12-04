package blendvpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvpd_128.s
var assemblyVblendvpd128 string

//go:embed stub_vblendvpd_128.go
var stubVblendvpd128 string

type VBLENDVPD128 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPD128() *VBLENDVPD128 {
	return &VBLENDVPD128{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		mask:  number.NewNamedUintParameter("mask", 128, 64, 16),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDVPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPD128) Name() string {
	return "VBLENDVPD (128 bit) "
}

func (v *VBLENDVPD128) Description() string {
	return "Blend packed double-precision floats using per-lane vector mask (sign bits)."
}

func (v *VBLENDVPD128) Stub() string {
	return stubVblendvpd128
}

func (v *VBLENDVPD128) Assembly() string {
	return assemblyVblendvpd128
}

func (v *VBLENDVPD128) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))
	mask := [2]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))

	ret := [2]float64{}

	vblendvpd128(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPD128 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDVPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
