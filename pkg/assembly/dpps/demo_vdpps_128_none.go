package dpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdpps_128_none.s
var assemblyVdpps128None string

//go:embed stub_vdpps_128_none.go
var stubVdpps128None string

type VDPPS128NONE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPS128NONE() *VDPPS128NONE {
	return &VDPPS128NONE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VDPPS128NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPS128NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPS128NONE) Name() string {
	return "VDPPS (128 bit) none"
}

func (v *VDPPS128NONE) Description() string {
	return "Dot product of packed singles with imm8=0x00 (no lanes selected)."
}

func (v *VDPPS128NONE) Stub() string {
	return stubVdpps128None
}

func (v *VDPPS128NONE) Assembly() string {
	return assemblyVdpps128None
}

func (v *VDPPS128NONE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vdpps128None(&vals1, &vals2, &ret)

	log.Printf("VDPPS128NONE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPS128NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
