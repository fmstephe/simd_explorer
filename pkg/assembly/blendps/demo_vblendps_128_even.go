package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_128_even.s
var assemblyVblendps128Even string

//go:embed stub_vblendps_128_even.go
var stubVblendps128Even string

type VBLENDPS128EVEN struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS128EVEN() *VBLENDPS128EVEN {
	return &VBLENDPS128EVEN{
		base:  number.NewNamedFloatParameter("base", 128, 32),
		blend: number.NewNamedFloatParameter("blend", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VBLENDPS128EVEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS128EVEN) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS128EVEN) Name() string {
	return "VBLENDPS (128 bit) even"
}

func (v *VBLENDPS128EVEN) Description() string {
	return "Blend 4 float32 lanes: imm=0x5 selects lanes 0,2 from blend; others from base."
}

func (v *VBLENDPS128EVEN) Stub() string {
	return stubVblendps128Even
}

func (v *VBLENDPS128EVEN) Assembly() string {
	return assemblyVblendps128Even
}

func (v *VBLENDPS128EVEN) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [4]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [4]float32{}

	vblendps128Even(&base, &blend, &ret)

	log.Printf("VBLENDPS128EVEN base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS128EVEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
