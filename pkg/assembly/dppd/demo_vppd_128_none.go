package dppd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdppd_128_none.s
var assemblyVdppd128None string

//go:embed stub_vdppd_128_none.go
var stubVdppd128None string

type VDPPD128NONE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDPPD128NONE() *VDPPD128NONE {
	return &VDPPD128NONE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VDPPD128NONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDPPD128NONE) Output() *number.Parameter {
	return v.ret
}

func (v *VDPPD128NONE) Name() string {
	return "VDPPD (128 bit) none"
}

func (v *VDPPD128NONE) Description() string {
	return "Dot product of packed doubles with imm8=0x00 (no lanes selected)."
}

func (v *VDPPD128NONE) Stub() string {
	return stubVdppd128None
}

func (v *VDPPD128NONE) Assembly() string {
	return assemblyVdppd128None
}

func (v *VDPPD128NONE) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))

	ret := [2]float64{}

	vdppd128None(&vals1, &vals2, &ret)

	log.Printf("VDPPD128NONE vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VDPPD128NONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
