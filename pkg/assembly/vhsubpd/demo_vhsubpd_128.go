package vhsubpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhsubpd_128.s
var assemblyVhsubpd128 string

//go:embed stub_vhsubpd_128.go
var stubVhsubpd128 string

type VHSUBPD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHSUBPD128() *VHSUBPD128 {
	return &VHSUBPD128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VHSUBPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHSUBPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VHSUBPD128) Name() string {
	return "VHSUBPD (128 bit) "
}

func (v *VHSUBPD128) Description() string {
	return "Horizontal subtract of packed double-precision floats. " +
		"Computes pairwise differences within each source (a0-a1) and (b0-b1), writes results to ret."
}

func (v *VHSUBPD128) Stub() string {
	return stubVhsubpd128
}

func (v *VHSUBPD128) Assembly() string {
	return assemblyVhsubpd128
}

func (v *VHSUBPD128) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vhsubpd128(&vals1, &vals2, &ret)

	log.Printf("VHSUBPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHSUBPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
