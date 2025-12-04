package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_256_all.s
var assemblyVblendps256All string

//go:embed stub_vblendps_256_all.go
var stubVblendps256All string

type VBLENDPS256ALL struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS256ALL() *VBLENDPS256ALL {
	return &VBLENDPS256ALL{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDPS256ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS256ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS256ALL) Name() string {
	return "VBLENDPS (256 bit) all"
}

func (v *VBLENDPS256ALL) Description() string {
	return "Blend 8 float32 lanes: imm=0xFF selects all lanes from blend."
}

func (v *VBLENDPS256ALL) Stub() string {
	return stubVblendps256All
}

func (v *VBLENDPS256ALL) Assembly() string {
	return assemblyVblendps256All
}

func (v *VBLENDPS256ALL) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [8]float32{}

	vblendps256All(&base, &blend, &ret)

	log.Printf("VBLENDPS256ALL base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS256ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
