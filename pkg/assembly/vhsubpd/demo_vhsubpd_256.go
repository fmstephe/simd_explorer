package vhsubpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhsubpd_256.s
var assemblyVhsubpd256 string

//go:embed stub_vhsubpd_256.go
var stubVhsubpd256 string

type VHSUBPD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHSUBPD256() *VHSUBPD256 {
	return &VHSUBPD256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VHSUBPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHSUBPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VHSUBPD256) Name() string {
	return "VHSUBPD (256 bit) "
}

func (v *VHSUBPD256) Description() string {
	return "Horizontal subtract of packed double-precision floats. " +
		"Computes pairwise differences within each 128-bit lane; results stored in ret."
}

func (v *VHSUBPD256) Stub() string {
	return stubVhsubpd256
}

func (v *VHSUBPD256) Assembly() string {
	return assemblyVhsubpd256
}

func (v *VHSUBPD256) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vhsubpd256(&vals1, &vals2, &ret)

	log.Printf("VHSUBPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHSUBPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
