package dpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdpps_128_all.s
var assemblyVdpps128All string

//go:embed stub_vdpps_128_all.go
var stubVdpps128All string

type VDPPS128ALL struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPS128ALL() *VDPPS128ALL {
	return &VDPPS128ALL{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VDPPS128ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPS128ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPS128ALL) Name() string {
	return "VDPPS (128 bit) all"
}

func (v *VDPPS128ALL) Description() string {
	return "Dot product of packed singles with imm8=0xFF (write result to all elements)."
}

func (v *VDPPS128ALL) Stub() string {
	return stubVdpps128All
}

func (v *VDPPS128ALL) Assembly() string {
	return assemblyVdpps128All
}

func (v *VDPPS128ALL) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vdpps128All(&vals1, &vals2, &ret)

	log.Printf("VDPPS128ALL vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPS128ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
