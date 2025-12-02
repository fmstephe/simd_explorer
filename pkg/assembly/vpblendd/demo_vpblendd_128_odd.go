package vpblendd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpblendd_128_odd.s
var assemblyVpblendd128Odd string

//go:embed stub_vpblendd_128_odd.go
var stubVpblendd128Odd string

type VPBLENDD128ODD struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVPBLENDD128ODD() *VPBLENDD128ODD {
	return &VPBLENDD128ODD{
		base:  number.NewNamedUintParameter("base", 128, 32, 16),
		blend: number.NewNamedUintParameter("blend", 128, 32, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPBLENDD128ODD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VPBLENDD128ODD) Output() *number.Parameter {
	return v.ret
}

func (v *VPBLENDD128ODD) Name() string {
	return "VPBLENDD (128 bit) odd"
}

func (v *VPBLENDD128ODD) Description() string {
	return "Blend 4 u32 dwords: imm=0xA selects odd lanes (1,3) from blend; others from base."
}

func (v *VPBLENDD128ODD) Stub() string {
	return stubVpblendd128Odd
}

func (v *VPBLENDD128ODD) Assembly() string {
	return assemblyVpblendd128Odd
}

func (v *VPBLENDD128ODD) Run() {
	base := [4]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	blend := [4]uint32{}
	copy(blend[:], number.ToUint32Slice(v.blend.FlatData()))

	ret := [4]uint32{}

	vpblendd128Odd(&base, &blend, &ret)

	log.Printf("VPBLENDD128ODD base %v blend %v ret %v", base, blend, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPBLENDD128ODD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
