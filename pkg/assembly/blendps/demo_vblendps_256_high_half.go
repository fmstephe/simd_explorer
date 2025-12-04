package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_256_high_half.s
var assemblyVblendps256High_half string

//go:embed stub_vblendps_256_high_half.go
var stubVblendps256High_half string

type VBLENDPS256HIGH_HALF struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS256HIGH_HALF() *VBLENDPS256HIGH_HALF {
	return &VBLENDPS256HIGH_HALF{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDPS256HIGH_HALF) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS256HIGH_HALF) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS256HIGH_HALF) Name() string {
	return "VBLENDPS (256 bit) high_half"
}

func (v *VBLENDPS256HIGH_HALF) Description() string {
	return "Blend 8 float32 lanes: imm=0xF0 selects lanes 4..7 from blend; others from base."
}

func (v *VBLENDPS256HIGH_HALF) Stub() string {
	return stubVblendps256High_half
}

func (v *VBLENDPS256HIGH_HALF) Assembly() string {
	return assemblyVblendps256High_half
}

func (v *VBLENDPS256HIGH_HALF) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [8]float32{}

	vblendps256High_half(&base, &blend, &ret)

	log.Printf("VBLENDPS256HIGH_HALF base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS256HIGH_HALF) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
