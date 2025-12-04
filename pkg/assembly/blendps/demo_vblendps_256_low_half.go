package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_256_low_half.s
var assemblyVblendps256Low_half string

//go:embed stub_vblendps_256_low_half.go
var stubVblendps256Low_half string

type VBLENDPS256LOW_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS256LOW_HALF() *VBLENDPS256LOW_HALF {
	return &VBLENDPS256LOW_HALF{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDPS256LOW_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS256LOW_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS256LOW_HALF) Name() string {
	return "VBLENDPS (256 bit) low_half"
}

func (v *VBLENDPS256LOW_HALF) Description() string {
	return "Blend 8 float32 lanes: imm=0x0F selects lanes 0..3 from blend; others from base."
}

func (v *VBLENDPS256LOW_HALF) Stub() string {
	return stubVblendps256Low_half
}

func (v *VBLENDPS256LOW_HALF) Assembly() string {
	return assemblyVblendps256Low_half
}

func (v *VBLENDPS256LOW_HALF) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [8]float32{}

	vblendps256Low_half(&base, &blend, &ret)

	log.Printf("VBLENDPS256LOW_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS256LOW_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
