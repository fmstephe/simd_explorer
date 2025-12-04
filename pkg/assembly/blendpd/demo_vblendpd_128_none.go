package blendpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendpd_128_none.s
var assemblyVblendpd128None string

//go:embed stub_vblendpd_128_none.go
var stubVblendpd128None string

type VBLENDPD128NONE struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPD128NONE() *VBLENDPD128NONE {
	return &VBLENDPD128NONE{
		base:  number.NewNamedFloatParameter("base", 128, 64),
		blend: number.NewNamedFloatParameter("blend", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VBLENDPD128NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPD128NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPD128NONE) Name() string {
	return "VBLENDPD (128 bit) none"
}

func (v *VBLENDPD128NONE) Description() string {
	return "Blend 2 float64 lanes: imm=0x0 selects all lanes from base."
}

func (v *VBLENDPD128NONE) Stub() string {
	return stubVblendpd128None
}

func (v *VBLENDPD128NONE) Assembly() string {
	return assemblyVblendpd128None
}

func (v *VBLENDPD128NONE) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	blend := [2]float64{}
	copy(blend[:], number.ToFloat64Slice(v.blend.FlatData()))

	ret := [2]float64{}

	vblendpd128None(&base, &blend, &ret)

	log.Printf("VBLENDPD128NONE base %v blend %v ret %v", base, blend, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPD128NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
