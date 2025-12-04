package blendvps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvps_256.s
var assemblyVblendvps256 string

//go:embed stub_vblendvps_256.go
var stubVblendvps256 string

type VBLENDVPS256 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPS256() *VBLENDVPS256 {
	return &VBLENDVPS256{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		mask:  number.NewNamedUintParameter("mask", 256, 32, 16),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDVPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPS256) Name() string {
	return "VBLENDVPS (256 bit) "
}

func (v *VBLENDVPS256) Description() string {
	return "Blend packed single-precision floats using per-lane vector mask (sign bits)."
}

func (v *VBLENDVPS256) Stub() string {
	return stubVblendvps256
}

func (v *VBLENDVPS256) Assembly() string {
	return assemblyVblendvps256
}

func (v *VBLENDVPS256) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))
	mask := [8]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))

	ret := [8]float32{}

	vblendvps256(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPS256 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDVPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
