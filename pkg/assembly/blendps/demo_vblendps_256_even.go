package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_256_even.s
var assemblyVblendps256Even string

//go:embed stub_vblendps_256_even.go
var stubVblendps256Even string

type VBLENDPS256EVEN struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS256EVEN() *VBLENDPS256EVEN {
	return &VBLENDPS256EVEN{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDPS256EVEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS256EVEN) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS256EVEN) Name() string {
	return "VBLENDPS (256 bit) even"
}

func (v *VBLENDPS256EVEN) Description() string {
	return "Blend 8 float32 lanes: imm=0x55 selects even lanes (0,2,4,6) from blend; others from base."
}

func (v *VBLENDPS256EVEN) Stub() string {
	return stubVblendps256Even
}

func (v *VBLENDPS256EVEN) Assembly() string {
	return assemblyVblendps256Even
}

func (v *VBLENDPS256EVEN) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [8]float32{}

	vblendps256Even(&base, &blend, &ret)

	log.Printf("VBLENDPS256EVEN base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS256EVEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
