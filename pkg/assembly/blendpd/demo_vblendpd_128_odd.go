package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_128_odd.s
var assemblyVblendpd128Odd string

//go:embed stub_vblendpd_128_odd.go
var stubVblendpd128Odd string

type VBLENDPD128ODD struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD128ODD() *VBLENDPD128ODD {
	return &VBLENDPD128ODD{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDPD128ODD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD128ODD) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD128ODD) Name() string {
	return "VBLENDPD (128 bit) odd"
}

func (v *VBLENDPD128ODD) Description() string {
	return "Blend 2 float64 lanes: imm=0x2 selects lane 1 from blend; lane 0 from base."
}

func (v *VBLENDPD128ODD) Stub() string {
	return stubVblendpd128Odd
}

func (v *VBLENDPD128ODD) Assembly() string {
	return assemblyVblendpd128Odd
}

func (v *VBLENDPD128ODD) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [2]float64{}

	vblendpd128Odd(&base, &blend, &ret)

	log.Printf("VBLENDPD128ODD base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD128ODD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
