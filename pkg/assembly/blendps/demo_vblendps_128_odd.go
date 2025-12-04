package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_128_odd.s
var assemblyVblendps128Odd string

//go:embed stub_vblendps_128_odd.go
var stubVblendps128Odd string

type VBLENDPS128ODD struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS128ODD() *VBLENDPS128ODD {
	return &VBLENDPS128ODD{
		base:  number.NewNamedFloatParameter("base", 128, 32),
		blend: number.NewNamedFloatParameter("blend", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VBLENDPS128ODD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS128ODD) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS128ODD) Name() string {
	return "VBLENDPS (128 bit) odd"
}

func (v *VBLENDPS128ODD) Description() string {
	return "Blend 4 float32 lanes: imm=0xA selects lanes 1,3 from blend; others from base."
}

func (v *VBLENDPS128ODD) Stub() string {
	return stubVblendps128Odd
}

func (v *VBLENDPS128ODD) Assembly() string {
	return assemblyVblendps128Odd
}

func (v *VBLENDPS128ODD) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [4]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [4]float32{}

	vblendps128Odd(&base, &blend, &ret)

	log.Printf("VBLENDPS128ODD base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS128ODD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
