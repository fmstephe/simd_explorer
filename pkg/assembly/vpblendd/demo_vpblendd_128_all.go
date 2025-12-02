package vpblendd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpblendd_128_all.s
var assemblyVpblendd128All string

//go:embed stub_vpblendd_128_all.go
var stubVpblendd128All string

type VPBLENDD128ALL struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVPBLENDD128ALL() *VPBLENDD128ALL {
	return &VPBLENDD128ALL{
		base:  number.NewNamedUintParameter("base", 128, 32, 16),
		blend: number.NewNamedUintParameter("blend", 128, 32, 16),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VPBLENDD128ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VPBLENDD128ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VPBLENDD128ALL) Name() string {
	return "VPBLENDD (128 bit) all"
}

func (v *VPBLENDD128ALL) Description() string {
	return "Blend 4 u32 dwords: imm=0xF selects all lanes from blend."
}

func (v *VPBLENDD128ALL) Stub() string {
	return stubVpblendd128All
}

func (v *VPBLENDD128ALL) Assembly() string {
	return assemblyVpblendd128All
}

func (v *VPBLENDD128ALL) Run() {
	base := [4]uint32{}
	copy(base[:], number.ToUint32Slice(v.base.FlatData()))
	blend := [4]uint32{}
	copy(blend[:], number.ToUint32Slice(v.blend.FlatData()))

	ret := [4]uint32{}

	vpblendd128All(&base, &blend, &ret)

	log.Printf("VPBLENDD128ALL base %v blend %v ret %v", base, blend, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPBLENDD128ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
