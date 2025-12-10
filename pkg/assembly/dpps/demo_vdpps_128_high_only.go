package dpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdpps_128_high_only.s
var assemblyVdpps128High_only string

//go:embed stub_vdpps_128_high_only.go
var stubVdpps128High_only string

type VDPPS128HIGH_ONLY struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPS128HIGH_ONLY() *VDPPS128HIGH_ONLY {
	return &VDPPS128HIGH_ONLY{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VDPPS128HIGH_ONLY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPS128HIGH_ONLY) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPS128HIGH_ONLY) Name() string {
	return "VDPPS (128 bit) high_only"
}

func (v *VDPPS128HIGH_ONLY) Description() string {
	return "Dot product of packed singles with imm8=0x8F (write high element only)."
}

func (v *VDPPS128HIGH_ONLY) Stub() string {
	return stubVdpps128High_only
}

func (v *VDPPS128HIGH_ONLY) Assembly() string {
	return assemblyVdpps128High_only
}

func (v *VDPPS128HIGH_ONLY) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vdpps128High_only(&vals1, &vals2, &ret)

	log.Printf("VDPPS128HIGH_ONLY vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPS128HIGH_ONLY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
