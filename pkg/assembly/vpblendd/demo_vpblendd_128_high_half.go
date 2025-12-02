package vpblendd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpblendd_128_high_half.s
var assemblyVpblendd128High_half string

//go:embed stub_vpblendd_128_high_half.go
var stubVpblendd128High_half string

type VPBLENDD128HIGH_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVPBLENDD128HIGH_HALF() *VPBLENDD128HIGH_HALF {
	return &VPBLENDD128HIGH_HALF{
		base:  number.NewNamedUintParameter("base", 128, 32, 16),
		blend: number.NewNamedUintParameter("blend", 128, 32, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPBLENDD128HIGH_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VPBLENDD128HIGH_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VPBLENDD128HIGH_HALF) Name() string {
	return "VPBLENDD (128 bit) high_half"
}

func (v *VPBLENDD128HIGH_HALF) Description() string {
	return "Blend 4 u32 dwords: imm=0xC selects high half (lanes 2,3) from blend; others from base."
}

func (v *VPBLENDD128HIGH_HALF) Stub() string {
	return stubVpblendd128High_half
}

func (v *VPBLENDD128HIGH_HALF) Assembly() string {
	return assemblyVpblendd128High_half
}

func (v *VPBLENDD128HIGH_HALF) Run() {
	base := [4]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	blend := [4]uint32{}
	copy(blend[:], number.ToUint32Slice(v.blend.FlatData()))

	ret := [4]uint32{}

	vpblendd128High_half(&base, &blend, &ret)

	log.Printf("VPBLENDD128HIGH_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPBLENDD128HIGH_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
