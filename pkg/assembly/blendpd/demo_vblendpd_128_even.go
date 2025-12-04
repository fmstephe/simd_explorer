package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_128_even.s
var assemblyVblendpd128Even string

//go:embed stub_vblendpd_128_even.go
var stubVblendpd128Even string

type VBLENDPD128EVEN struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD128EVEN() *VBLENDPD128EVEN {
	return &VBLENDPD128EVEN{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDPD128EVEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD128EVEN) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD128EVEN) Name() string {
	return "VBLENDPD (128 bit) even"
}

func (v *VBLENDPD128EVEN) Description() string {
	return "Blend 2 float64 lanes: imm=0x1 selects lane 0 from blend; lane 1 from base."
}

func (v *VBLENDPD128EVEN) Stub() string {
	return stubVblendpd128Even
}

func (v *VBLENDPD128EVEN) Assembly() string {
	return assemblyVblendpd128Even
}

func (v *VBLENDPD128EVEN) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [2]float64{}

	vblendpd128Even(&base, &blend, &ret)

	log.Printf("VBLENDPD128EVEN base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD128EVEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
