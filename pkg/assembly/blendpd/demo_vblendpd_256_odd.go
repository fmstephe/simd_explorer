package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_256_odd.s
var assemblyVblendpd256Odd string

//go:embed stub_vblendpd_256_odd.go
var stubVblendpd256Odd string

type VBLENDPD256ODD struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD256ODD() *VBLENDPD256ODD {
	return &VBLENDPD256ODD{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDPD256ODD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD256ODD) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD256ODD) Name() string {
	return "VBLENDPD (256 bit) odd"
}

func (v *VBLENDPD256ODD) Description() string {
	return "Blend 4 float64 lanes: imm=0xA selects lanes 1,3 from blend; others from base."
}

func (v *VBLENDPD256ODD) Stub() string {
	return stubVblendpd256Odd
}

func (v *VBLENDPD256ODD) Assembly() string {
	return assemblyVblendpd256Odd
}

func (v *VBLENDPD256ODD) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [4]float64{}

	vblendpd256Odd(&base, &blend, &ret)

	log.Printf("VBLENDPD256ODD base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD256ODD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
