package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_256_odd.s
var assemblyVblendps256Odd string

//go:embed stub_vblendps_256_odd.go
var stubVblendps256Odd string

type VBLENDPS256ODD struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS256ODD() *VBLENDPS256ODD {
	return &VBLENDPS256ODD{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDPS256ODD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS256ODD) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS256ODD) Name() string {
	return "VBLENDPS (256 bit) odd"
}

func (v *VBLENDPS256ODD) Description() string {
	return "Blend 8 float32 lanes: imm=0xAA selects odd lanes (1,3,5,7) from blend; others from base."
}

func (v *VBLENDPS256ODD) Stub() string {
	return stubVblendps256Odd
}

func (v *VBLENDPS256ODD) Assembly() string {
	return assemblyVblendps256Odd
}

func (v *VBLENDPS256ODD) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [8]float32{}

	vblendps256Odd(&base, &blend, &ret)

	log.Printf("VBLENDPS256ODD base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS256ODD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
