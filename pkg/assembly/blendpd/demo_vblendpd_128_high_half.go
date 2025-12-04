package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_128_high_half.s
var assemblyVblendpd128High_half string

//go:embed stub_vblendpd_128_high_half.go
var stubVblendpd128High_half string

type VBLENDPD128HIGH_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD128HIGH_HALF() *VBLENDPD128HIGH_HALF {
	return &VBLENDPD128HIGH_HALF{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDPD128HIGH_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD128HIGH_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD128HIGH_HALF) Name() string {
	return "VBLENDPD (128 bit) high_half"
}

func (v *VBLENDPD128HIGH_HALF) Description() string {
	return "Blend 2 float64 lanes: imm=0x2 selects lane 1 from blend; lane 0 from base."
}

func (v *VBLENDPD128HIGH_HALF) Stub() string {
	return stubVblendpd128High_half
}

func (v *VBLENDPD128HIGH_HALF) Assembly() string {
	return assemblyVblendpd128High_half
}

func (v *VBLENDPD128HIGH_HALF) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [2]float64{}

	vblendpd128High_half(&base, &blend, &ret)

	log.Printf("VBLENDPD128HIGH_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD128HIGH_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
