package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_128_all.s
var assemblyVblendpd128All string

//go:embed stub_vblendpd_128_all.go
var stubVblendpd128All string

type VBLENDPD128ALL struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD128ALL() *VBLENDPD128ALL {
	return &VBLENDPD128ALL{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDPD128ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD128ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD128ALL) Name() string {
	return "VBLENDPD (128 bit) all"
}

func (v *VBLENDPD128ALL) Description() string {
	return "Blend 2 float64 lanes: imm=0x3 selects both lanes from blend."
}

func (v *VBLENDPD128ALL) Stub() string {
	return stubVblendpd128All
}

func (v *VBLENDPD128ALL) Assembly() string {
	return assemblyVblendpd128All
}

func (v *VBLENDPD128ALL) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [2]float64{}

	vblendpd128All(&base, &blend, &ret)

	log.Printf("VBLENDPD128ALL base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD128ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
