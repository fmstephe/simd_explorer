package vhaddpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhaddpd_128.s
var assemblyVhaddpd128 string

//go:embed stub_vhaddpd_128.go
var stubVhaddpd128 string

type VHADDPD128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHADDPD128() *VHADDPD128 {
	return &VHADDPD128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 64),
		vals2: number.NewNamedFloatParameter("vals2", 128, 64),
		ret:   number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VHADDPD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHADDPD128) Output() *number.Parameter {
	return v.ret
}

func (v *VHADDPD128) Name() string {
	return "VHADDPD (128 bit) "
}

func (v *VHADDPD128) Description() string {
	return "Horizontal add of packed double-precision floats. " +
		"Computes pairwise sums within each source (a0+a1) and (b0+b1), writes results to ret."
}

func (v *VHADDPD128) Stub() string {
	return stubVhaddpd128
}

func (v *VHADDPD128) Assembly() string {
	return assemblyVhaddpd128
}

func (v *VHADDPD128) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [2]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vhaddpd128(&vals1, &vals2, &ret)

	log.Printf("VHADDPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHADDPD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
