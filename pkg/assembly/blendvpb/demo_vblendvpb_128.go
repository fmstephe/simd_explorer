package blendvpb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvpb_128.s
var assemblyVblendvpb128 string

//go:embed stub_vblendvpb_128.go
var stubVblendvpb128 string

type VBLENDVPB128 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPB128() *VBLENDVPB128 {
	return &VBLENDVPB128{
		base:  number.NewNamedUintParameter("base", 128, 8, 10),
		blend: number.NewNamedUintParameter("blend", 128, 8, 10),
		mask:  number.NewNamedUintParameter("mask", 128, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VBLENDVPB128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPB128) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPB128) Name() string {
	return "VBLENDVPB (128 bit) "
}

func (v *VBLENDVPB128) Description() string {
	return "Blend 16 u8 bytes using per-byte mask sign bits (VPBLENDVB)."
}

func (v *VBLENDVPB128) Stub() string {
	return stubVblendvpb128
}

func (v *VBLENDVPB128) Assembly() string {
	return assemblyVblendvpb128
}

func (v *VBLENDVPB128) Run() {
	base := [16]uint8{}
	copy(base[:], v.base.FlatData())
	blend := [16]uint8{}
	copy(blend[:], v.blend.FlatData())
	mask := [16]uint8{}
	copy(mask[:], v.mask.FlatData())

	ret := [16]uint8{}

	vblendvpb128(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPB128 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := ret[:]
	v.ret.SetData(out)
}

func (v *VBLENDVPB128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
