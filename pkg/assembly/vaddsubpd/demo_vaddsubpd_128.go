package vaddsubpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddsubpd_128.s
var assemblyVaddsubpd128 string

//go:embed stub_vaddsubpd_128.go
var stubVaddsubpd128 string

type VADDSUBPD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDSUBPD128() *VADDSUBPD128 {
	return &VADDSUBPD128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VADDSUBPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDSUBPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VADDSUBPD128) Name() string {
	return "VADDSUBPD (128 bit) "
}

func (v *VADDSUBPD128) Description() string {
	return "Add-subtract packed double-precision floats. " +
		"Even elements are added (a0+b0, a2+b2), odd elements are subtracted (a1-b1, a3-b3) and stored in ret."
}

func (v *VADDSUBPD128) Stub() string {
	return stubVaddsubpd128
}

func (v *VADDSUBPD128) Assembly() string {
	return assemblyVaddsubpd128
}

func (v *VADDSUBPD128) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vaddsubpd128(&vals1, &vals2, &ret)

	log.Printf("VADDSUBPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VADDSUBPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
