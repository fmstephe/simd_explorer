package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_256_none.s
var assemblyVblendpd256None string

//go:embed stub_vblendpd_256_none.go
var stubVblendpd256None string

type VBLENDPD256NONE struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD256NONE() *VBLENDPD256NONE {
	return &VBLENDPD256NONE{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDPD256NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD256NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD256NONE) Name() string {
	return "VBLENDPD (256 bit) none"
}

func (v *VBLENDPD256NONE) Description() string {
	return "Blend 4 float64 lanes: imm=0x0 selects all lanes from base."
}

func (v *VBLENDPD256NONE) Stub() string {
	return stubVblendpd256None
}

func (v *VBLENDPD256NONE) Assembly() string {
	return assemblyVblendpd256None
}

func (v *VBLENDPD256NONE) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [4]float64{}

	vblendpd256None(&base, &blend, &ret)

	log.Printf("VBLENDPD256NONE base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD256NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
