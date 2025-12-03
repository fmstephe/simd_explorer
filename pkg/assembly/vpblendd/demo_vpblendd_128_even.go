package vpblendd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpblendd_128_even.s
var assemblyVpblendd128Even string

//go:embed stub_vpblendd_128_even.go
var stubVpblendd128Even string

type VPBLENDD128EVEN struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVPBLENDD128EVEN() *VPBLENDD128EVEN {
	return &VPBLENDD128EVEN{
		base:  number.NewNamedUintParameter("base", 128, 32, 10),
		blend: number.NewNamedUintParameter("blend", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPBLENDD128EVEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VPBLENDD128EVEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPBLENDD128EVEN) Name() string {
	return "VPBLENDD (128 bit) even"
}

func (v *VPBLENDD128EVEN) Description() string {
	return "Blend 4 u32 dwords: imm=0x5 selects even lanes (0,2) from blend; others from base."
}

func (v *VPBLENDD128EVEN) Stub() string {
	return stubVpblendd128Even
}

func (v *VPBLENDD128EVEN) Assembly() string {
	return assemblyVpblendd128Even
}

func (v *VPBLENDD128EVEN) Run() {
	base := [4]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	blend := [4]uint32{}
	copy(blend[:], number.ToUint32Slice(v.blend.FlatData()))

	ret := [4]uint32{}

	vpblendd128Even(&base, &blend, &ret)

	log.Printf("VPBLENDD128EVEN base %v blend %v ret %v", base, blend, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPBLENDD128EVEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
