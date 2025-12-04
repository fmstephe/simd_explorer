package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_128_none.s
var assemblyVblendps128None string

//go:embed stub_vblendps_128_none.go
var stubVblendps128None string

type VBLENDPS128NONE struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS128NONE() *VBLENDPS128NONE {
	return &VBLENDPS128NONE{
		base:  number.NewNamedFloatParameter("base", 128, 32),
		blend: number.NewNamedFloatParameter("blend", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VBLENDPS128NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS128NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS128NONE) Name() string {
	return "VBLENDPS (128 bit) none"
}

func (v *VBLENDPS128NONE) Description() string {
	return "Blend 4 float32 lanes: imm=0x0 selects all lanes from base."
}

func (v *VBLENDPS128NONE) Stub() string {
	return stubVblendps128None
}

func (v *VBLENDPS128NONE) Assembly() string {
	return assemblyVblendps128None
}

func (v *VBLENDPS128NONE) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [4]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [4]float32{}

	vblendps128None(&base, &blend, &ret)

	log.Printf("VBLENDPS128NONE base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS128NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
