package blendvpb

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendvpb_256.s
var assemblyVblendvpb256 string

//go:embed stub_vblendvpb_256.go
var stubVblendvpb256 string

type VBLENDVPB256 struct {
	base  *number.Parameter
	blend *number.Parameter
	mask  *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDVPB256() *VBLENDVPB256 {
	return &VBLENDVPB256{
		base:  number.NewNamedUintParameter("base", 256, 8, 10),
		blend: number.NewNamedUintParameter("blend", 256, 8, 10),
		mask:  number.NewNamedUintParameter("mask", 256, 8, 16),
		ret:   number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VBLENDVPB256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
		v.mask,
	}
}

func (v *VBLENDVPB256) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDVPB256) Name() string {
	return "VBLENDVPB (256 bit) "
}

func (v *VBLENDVPB256) Description() string {
	return "Blend 32 u8 bytes using per-byte mask sign bits (VPBLENDVB)."
}

func (v *VBLENDVPB256) Stub() string {
	return stubVblendvpb256
}

func (v *VBLENDVPB256) Assembly() string {
	return assemblyVblendvpb256
}

func (v *VBLENDVPB256) Run() {
	base := [32]uint8{}
	copy(base[:], v.base.FlatData())
	blend := [32]uint8{}
	copy(blend[:], v.blend.FlatData())
	mask := [32]uint8{}
	copy(mask[:], v.mask.FlatData())

	ret := [32]uint8{}

	vblendvpb256(&base, &blend, &mask, &ret)

	log.Printf("VBLENDVPB256 base %v blend %v mask %v ret %v", base, blend, mask, ret)

	out := ret[:]
	v.ret.SetData(out)
}

func (v *VBLENDVPB256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
