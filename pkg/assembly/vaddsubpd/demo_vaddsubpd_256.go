package vaddsubpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddsubpd_256.s
var assemblyVaddsubpd256 string

//go:embed stub_vaddsubpd_256.go
var stubVaddsubpd256 string

type VADDSUBPD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDSUBPD256() *VADDSUBPD256 {
	return &VADDSUBPD256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VADDSUBPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDSUBPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VADDSUBPD256) Name() string {
	return "VADDSUBPD (256 bit) "
}

func (v *VADDSUBPD256) Description() string {
	return "Add-subtract packed double-precision floats. " +
		"Even elements are added and odd elements subtracted within each 128-bit lane; results stored in ret."
}

func (v *VADDSUBPD256) Stub() string {
	return stubVaddsubpd256
}

func (v *VADDSUBPD256) Assembly() string {
	return assemblyVaddsubpd256
}

func (v *VADDSUBPD256) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vaddsubpd256(&vals1, &vals2, &ret)

	log.Printf("VADDSUBPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VADDSUBPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
