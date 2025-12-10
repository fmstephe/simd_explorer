package dppd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdppd_128_all.s
var assemblyVdppd128All string

//go:embed stub_vdppd_128_all.go
var stubVdppd128All string

type VDPPD128ALL struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPD128ALL() *VDPPD128ALL {
	return &VDPPD128ALL{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VDPPD128ALL) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPD128ALL) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPD128ALL) Name() string {
	return "VDPPD (128 bit) all"
}

func (v *VDPPD128ALL) Description() string {
	return "Dot product of packed doubles with imm8=0x33 (write both elements)."
}

func (v *VDPPD128ALL) Stub() string {
	return stubVdppd128All
}

func (v *VDPPD128ALL) Assembly() string {
	return assemblyVdppd128All
}

func (v *VDPPD128ALL) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))

	ret := [2]float64{}

	vdppd128All(&vals1, &vals2, &ret)

	log.Printf("VDPPD128ALL vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPD128ALL) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
