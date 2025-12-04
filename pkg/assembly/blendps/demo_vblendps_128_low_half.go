package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_128_low_half.s
var assemblyVblendps128Low_half string

//go:embed stub_vblendps_128_low_half.go
var stubVblendps128Low_half string

type VBLENDPS128LOW_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS128LOW_HALF() *VBLENDPS128LOW_HALF {
	return &VBLENDPS128LOW_HALF{
		base:  number.NewNamedFloatParameter("base", 128, 32),
		blend: number.NewNamedFloatParameter("blend", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VBLENDPS128LOW_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS128LOW_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS128LOW_HALF) Name() string {
	return "VBLENDPS (128 bit) low_half"
}

func (v *VBLENDPS128LOW_HALF) Description() string {
	return "Blend 4 float32 lanes: imm=0x3 selects lanes 0,1 from blend; others from base."
}

func (v *VBLENDPS128LOW_HALF) Stub() string {
	return stubVblendps128Low_half
}

func (v *VBLENDPS128LOW_HALF) Assembly() string {
	return assemblyVblendps128Low_half
}

func (v *VBLENDPS128LOW_HALF) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [4]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [4]float32{}

	vblendps128Low_half(&base, &blend, &ret)

	log.Printf("VBLENDPS128LOW_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS128LOW_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
