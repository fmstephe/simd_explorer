package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_128_all.s
var assemblyVblendps128All string

//go:embed stub_vblendps_128_all.go
var stubVblendps128All string

type VBLENDPS128ALL struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS128ALL() *VBLENDPS128ALL {
	return &VBLENDPS128ALL{
		base:  number.NewNamedFloatParameter("base", 128, 32),
		blend: number.NewNamedFloatParameter("blend", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VBLENDPS128ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS128ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS128ALL) Name() string {
	return "VBLENDPS (128 bit) all"
}

func (v *VBLENDPS128ALL) Description() string {
	return "Blend 4 float32 lanes: imm=0xF selects all lanes from blend."
}

func (v *VBLENDPS128ALL) Stub() string {
	return stubVblendps128All
}

func (v *VBLENDPS128ALL) Assembly() string {
	return assemblyVblendps128All
}

func (v *VBLENDPS128ALL) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [4]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [4]float32{}

	vblendps128All(&base, &blend, &ret)

	log.Printf("VBLENDPS128ALL base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS128ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
