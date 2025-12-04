package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_256_high_half.s
var assemblyVblendpd256High_half string

//go:embed stub_vblendpd_256_high_half.go
var stubVblendpd256High_half string

type VBLENDPD256HIGH_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD256HIGH_HALF() *VBLENDPD256HIGH_HALF {
	return &VBLENDPD256HIGH_HALF{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDPD256HIGH_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD256HIGH_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD256HIGH_HALF) Name() string {
	return "VBLENDPD (256 bit) high_half"
}

func (v *VBLENDPD256HIGH_HALF) Description() string {
	return "Blend 4 float64 lanes: imm=0xF0 selects lanes 2,3 from blend; others from base."
}

func (v *VBLENDPD256HIGH_HALF) Stub() string {
	return stubVblendpd256High_half
}

func (v *VBLENDPD256HIGH_HALF) Assembly() string {
	return assemblyVblendpd256High_half
}

func (v *VBLENDPD256HIGH_HALF) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [4]float64{}

	vblendpd256High_half(&base, &blend, &ret)

	log.Printf("VBLENDPD256HIGH_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD256HIGH_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
