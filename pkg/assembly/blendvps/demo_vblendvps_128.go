package blendvps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvps_128.s
var assemblyVblendvps128 string

//go:embed stub_vblendvps_128.go
var stubVblendvps128 string

type VBLENDVPS128 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPS128() *VBLENDVPS128 {
	return &VBLENDVPS128{
		base:  number.NewNamedFloatParameter("base", 128, 32),
		blend: number.NewNamedFloatParameter("blend", 128, 32),
		mask:  number.NewNamedUintParameter("mask", 128, 32, 16),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VBLENDVPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPS128) Name() string {
	return "VBLENDVPS (128 bit) "
}

func (v *VBLENDVPS128) Description() string {
	return "Blend packed single-precision floats using per-lane vector mask (sign bits)."
}

func (v *VBLENDVPS128) Stub() string {
	return stubVblendvps128
}

func (v *VBLENDVPS128) Assembly() string {
	return assemblyVblendvps128
}

func (v *VBLENDVPS128) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [4]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))

	ret := [4]float32{}

	vblendvps128(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPS128 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDVPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
