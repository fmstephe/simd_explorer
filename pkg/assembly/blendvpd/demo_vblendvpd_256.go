package blendvpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvpd_256.s
var assemblyVblendvpd256 string

//go:embed stub_vblendvpd_256.go
var stubVblendvpd256 string

type VBLENDVPD256 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPD256() *VBLENDVPD256 {
	return &VBLENDVPD256{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		mask:  number.NewNamedUintParameter("mask", 256, 64, 16),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDVPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPD256) Name() string {
	return "VBLENDVPD (256 bit) "
}

func (v *VBLENDVPD256) Description() string {
	return "Blend packed double-precision floats using per-lane vector mask (sign bits)."
}

func (v *VBLENDVPD256) Stub() string {
	return stubVblendvpd256
}

func (v *VBLENDVPD256) Assembly() string {
	return assemblyVblendvpd256
}

func (v *VBLENDVPD256) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))
	mask := [4]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))

	ret := [4]float64{}

	vblendvpd256(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPD256 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDVPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
