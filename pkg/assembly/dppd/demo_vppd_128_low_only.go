package dppd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdppd_128_low_only.s
var assemblyVdppd128Low_only string

//go:embed stub_vdppd_128_low_only.go
var stubVdppd128Low_only string

type VDPPD128LOW_ONLY struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPD128LOW_ONLY() *VDPPD128LOW_ONLY {
	return &VDPPD128LOW_ONLY{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VDPPD128LOW_ONLY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPD128LOW_ONLY) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPD128LOW_ONLY) Name() string {
	return "VDPPD (128 bit) low_only"
}

func (v *VDPPD128LOW_ONLY) Description() string {
	return "Dot product of packed doubles with imm8=0x13 (write low element only)."
}

func (v *VDPPD128LOW_ONLY) Stub() string {
	return stubVdppd128Low_only
}

func (v *VDPPD128LOW_ONLY) Assembly() string {
	return assemblyVdppd128Low_only
}

func (v *VDPPD128LOW_ONLY) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))

	ret := [2]float64{}

	vdppd128Low_only(&vals1, &vals2, &ret)

	log.Printf("VDPPD128LOW_ONLY vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPD128LOW_ONLY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
