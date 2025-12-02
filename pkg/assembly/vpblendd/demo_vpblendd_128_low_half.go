package vpblendd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpblendd_128_low_half.s
var assemblyVpblendd128Low_half string

//go:embed stub_vpblendd_128_low_half.go
var stubVpblendd128Low_half string

type VPBLENDD128LOW_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVPBLENDD128LOW_HALF() *VPBLENDD128LOW_HALF {
	return &VPBLENDD128LOW_HALF{
		base:  number.NewNamedUintParameter("base", 128, 32, 16),
		blend: number.NewNamedUintParameter("blend", 128, 32, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPBLENDD128LOW_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VPBLENDD128LOW_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VPBLENDD128LOW_HALF) Name() string {
	return "VPBLENDD (128 bit) low_half"
}

func (v *VPBLENDD128LOW_HALF) Description() string {
	return "Blend 4 u32 dwords: imm=0x3 selects low half (lanes 0,1) from blend; others from base."
}

func (v *VPBLENDD128LOW_HALF) Stub() string {
	return stubVpblendd128Low_half
}

func (v *VPBLENDD128LOW_HALF) Assembly() string {
	return assemblyVpblendd128Low_half
}

func (v *VPBLENDD128LOW_HALF) Run() {
	base := [4]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	blend := [4]uint32{}
	copy(blend[:], number.ToUint32Slice(v.blend.FlatData()))

	ret := [4]uint32{}

	vpblendd128Low_half(&base, &blend, &ret)

	log.Printf("VPBLENDD128LOW_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPBLENDD128LOW_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
