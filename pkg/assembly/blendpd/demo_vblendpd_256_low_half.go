package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_256_low_half.s
var assemblyVblendpd256Low_half string

//go:embed stub_vblendpd_256_low_half.go
var stubVblendpd256Low_half string

type VBLENDPD256LOW_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD256LOW_HALF() *VBLENDPD256LOW_HALF {
	return &VBLENDPD256LOW_HALF{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDPD256LOW_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD256LOW_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD256LOW_HALF) Name() string {
	return "VBLENDPD (256 bit) low_half"
}

func (v *VBLENDPD256LOW_HALF) Description() string {
	return "Blend 4 float64 lanes: imm=0x3 selects lanes 0,1 from blend; others from base."
}

func (v *VBLENDPD256LOW_HALF) Stub() string {
	return stubVblendpd256Low_half
}

func (v *VBLENDPD256LOW_HALF) Assembly() string {
	return assemblyVblendpd256Low_half
}

func (v *VBLENDPD256LOW_HALF) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [4]float64{}

	vblendpd256Low_half(&base, &blend, &ret)

	log.Printf("VBLENDPD256LOW_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD256LOW_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
