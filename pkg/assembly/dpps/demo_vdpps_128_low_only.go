package dpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdpps_128_low_only.s
var assemblyVdpps128Low_only string

//go:embed stub_vdpps_128_low_only.go
var stubVdpps128Low_only string

type VDPPS128LOW_ONLY struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPS128LOW_ONLY() *VDPPS128LOW_ONLY {
	return &VDPPS128LOW_ONLY{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VDPPS128LOW_ONLY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPS128LOW_ONLY) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPS128LOW_ONLY) Name() string {
	return "VDPPS (128 bit) low_only"
}

func (v *VDPPS128LOW_ONLY) Description() string {
	return "Dot product of packed singles with imm8=0x1F (write low element only)."
}

func (v *VDPPS128LOW_ONLY) Stub() string {
	return stubVdpps128Low_only
}

func (v *VDPPS128LOW_ONLY) Assembly() string {
	return assemblyVdpps128Low_only
}

func (v *VDPPS128LOW_ONLY) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vdpps128Low_only(&vals1, &vals2, &ret)

	log.Printf("VDPPS128LOW_ONLY vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPS128LOW_ONLY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
