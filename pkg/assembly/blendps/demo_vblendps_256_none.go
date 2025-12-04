package blendps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vblendps_256_none.s
var assemblyVblendps256None string

//go:embed stub_vblendps_256_none.go
var stubVblendps256None string

type VBLENDPS256NONE struct {
	base  *number.Parameter
	blend *number.Parameter
	ret   *number.Parameter
}

func NewVBLENDPS256NONE() *VBLENDPS256NONE {
	return &VBLENDPS256NONE{
		base:  number.NewNamedFloatParameter("base", 256, 32),
		blend: number.NewNamedFloatParameter("blend", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VBLENDPS256NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.blend,
	}
}

func (v *VBLENDPS256NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VBLENDPS256NONE) Name() string {
	return "VBLENDPS (256 bit) none"
}

func (v *VBLENDPS256NONE) Description() string {
	return "Blend 8 float32 lanes: imm=0x00 selects all lanes from base."
}

func (v *VBLENDPS256NONE) Stub() string {
	return stubVblendps256None
}

func (v *VBLENDPS256NONE) Assembly() string {
	return assemblyVblendps256None
}

func (v *VBLENDPS256NONE) Run() {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	blend := [8]float32{}
	copy(blend[:], number.ToFloat32Slice(v.blend.FlatData()))

	ret := [8]float32{}

	vblendps256None(&base, &blend, &ret)

	log.Printf("VBLENDPS256NONE base %v blend %v ret %v", base, blend, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VBLENDPS256NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
