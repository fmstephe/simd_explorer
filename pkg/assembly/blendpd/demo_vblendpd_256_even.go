package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_256_even.s
var assemblyVblendpd256Even string

//go:embed stub_vblendpd_256_even.go
var stubVblendpd256Even string

type VBLENDPD256EVEN struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD256EVEN() *VBLENDPD256EVEN {
	return &VBLENDPD256EVEN{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDPD256EVEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD256EVEN) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD256EVEN) Name() string {
	return "VBLENDPD (256 bit) even"
}

func (v *VBLENDPD256EVEN) Description() string {
	return "Blend 4 float64 lanes: imm=0x5 selects lanes 0,2 from blend; others from base."
}

func (v *VBLENDPD256EVEN) Stub() string {
	return stubVblendpd256Even
}

func (v *VBLENDPD256EVEN) Assembly() string {
	return assemblyVblendpd256Even
}

func (v *VBLENDPD256EVEN) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [4]float64{}

	vblendpd256Even(&base, &blend, &ret)

	log.Printf("VBLENDPD256EVEN base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD256EVEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
