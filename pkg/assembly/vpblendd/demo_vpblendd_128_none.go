package vpblendd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpblendd_128_none.s
var assemblyVpblendd128None string

//go:embed stub_vpblendd_128_none.go
var stubVpblendd128None string

type VPBLENDD128NONE struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVPBLENDD128NONE() *VPBLENDD128NONE {
	return &VPBLENDD128NONE{
		base:  number.NewNamedUintParameter("base", 128, 32, 10),
		blend: number.NewNamedUintParameter("blend", 128, 32, 10),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPBLENDD128NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VPBLENDD128NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPBLENDD128NONE) Name() string {
	return "VPBLENDD (128 bit) none"
}

func (v *VPBLENDD128NONE) Description() string {
	return "Blend 4 u32 dwords: imm=0x0 selects all lanes from base."
}

func (v *VPBLENDD128NONE) Stub() string {
	return stubVpblendd128None
}

func (v *VPBLENDD128NONE) Assembly() string {
	return assemblyVpblendd128None
}

func (v *VPBLENDD128NONE) Run() {
	base := [4]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	blend := [4]uint32{}
	copy(blend[:], number.ToUint32Slice(v.blend.FlatData()))

	ret := [4]uint32{}

	vpblendd128None(&base, &blend, &ret)

	log.Printf("VPBLENDD128NONE base %v blend %v ret %v", base, blend, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPBLENDD128NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
