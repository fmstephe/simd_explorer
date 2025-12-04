package blendvpw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvpw_128.s
var assemblyVblendvpw128 string

//go:embed stub_vblendvpw_128.go
var stubVblendvpw128 string

type VBLENDVPW128 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPW128() *VBLENDVPW128 {
	return &VBLENDVPW128{
		base:  number.NewNamedUintParameter("base", 128, 16, 10),
		blend: number.NewNamedUintParameter("blend", 128, 16, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 16, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VBLENDVPW128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPW128) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPW128) Name() string {
	return "VBLENDVPW (128 bit) "
}

func (v *VBLENDVPW128) Description() string {
	return "Blend 8 u16 words using per-byte mask sign bits (VPBLENDVB)."
}

func (v *VBLENDVPW128) Stub() string {
	return stubVblendvpw128
}

func (v *VBLENDVPW128) Assembly() string {
	return assemblyVblendvpw128
}

func (v *VBLENDVPW128) Run() {
	base := [8]uint16{}
	copy(base[:], number.ToUint16Slice(v.base.FlatData()))
	blend := [8]uint16{}
	copy(blend[:], number.ToUint16Slice(v.blend.FlatData()))
	mask := [8]uint16{}
	copy(mask[:], number.ToUint16Slice(v.mask.FlatData()))

	ret := [8]uint16{}

	vblendvpw128(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPW128 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDVPW128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
