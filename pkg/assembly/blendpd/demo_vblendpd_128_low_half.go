package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_128_low_half.s
var assemblyVblendpd128Low_half string

//go:embed stub_vblendpd_128_low_half.go
var stubVblendpd128Low_half string

type VBLENDPD128LOW_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD128LOW_HALF() *VBLENDPD128LOW_HALF {
	return &VBLENDPD128LOW_HALF{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDPD128LOW_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD128LOW_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD128LOW_HALF) Name() string {
	return "VBLENDPD (128 bit) low_half"
}

func (v *VBLENDPD128LOW_HALF) Description() string {
	return "Blend 2 float64 lanes: imm=0x1 selects lane 0 from blend; lane 1 from base."
}

func (v *VBLENDPD128LOW_HALF) Stub() string {
	return stubVblendpd128Low_half
}

func (v *VBLENDPD128LOW_HALF) Assembly() string {
	return assemblyVblendpd128Low_half
}

func (v *VBLENDPD128LOW_HALF) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [2]float64{}

	vblendpd128Low_half(&base, &blend, &ret)

	log.Printf("VBLENDPD128LOW_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD128LOW_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
