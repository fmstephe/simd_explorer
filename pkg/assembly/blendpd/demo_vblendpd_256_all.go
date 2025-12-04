package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_256_all.s
var assemblyVblendpd256All string

//go:embed stub_vblendpd_256_all.go
var stubVblendpd256All string

type VBLENDPD256ALL struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD256ALL() *VBLENDPD256ALL {
	return &VBLENDPD256ALL{
		base:  number.NewNamedFloatParameter("base", 256, 64),
		blend: number.NewNamedFloatParameter("blend", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VBLENDPD256ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD256ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD256ALL) Name() string {
	return "VBLENDPD (256 bit) all"
}

func (v *VBLENDPD256ALL) Description() string {
	return "Blend 4 float64 lanes: imm=0xF selects all lanes from blend."
}

func (v *VBLENDPD256ALL) Stub() string {
	return stubVblendpd256All
}

func (v *VBLENDPD256ALL) Assembly() string {
	return assemblyVblendpd256All
}

func (v *VBLENDPD256ALL) Run() {
	base := [4]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [4]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [4]float64{}

	vblendpd256All(&base, &blend, &ret)

	log.Printf("VBLENDPD256ALL base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD256ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
