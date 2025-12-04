package blendvpw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvpw_256.s
var assemblyVblendvpw256 string

//go:embed stub_vblendvpw_256.go
var stubVblendvpw256 string

type VBLENDVPW256 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPW256() *VBLENDVPW256 {
	return &VBLENDVPW256{
		base:  number.NewNamedUintParameter("base", 256, 16, 10),
		blend: number.NewNamedUintParameter("blend", 256, 16, 10),
		mask:  number.NewNamedUintParameter("mask", 256, 16, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VBLENDVPW256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPW256) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPW256) Name() string {
	return "VBLENDVPW (256 bit) "
}

func (v *VBLENDVPW256) Description() string {
	return "Blend 16 u16 words using per-byte mask sign bits (VPBLENDVB)."
}

func (v *VBLENDVPW256) Stub() string {
	return stubVblendvpw256
}

func (v *VBLENDVPW256) Assembly() string {
	return assemblyVblendvpw256
}

func (v *VBLENDVPW256) Run() {
	base := [16]uint16{}
	copy(base[:], number.ToUint16Slice(v.base.FlatData()))
	blend := [16]uint16{}
	copy(blend[:], number.ToUint16Slice(v.blend.FlatData()))
	mask := [16]uint16{}
	copy(mask[:], number.ToUint16Slice(v.mask.FlatData()))

	ret := [16]uint16{}

	vblendvpw256(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPW256 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDVPW256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
