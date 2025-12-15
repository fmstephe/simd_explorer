package vhaddpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhaddpd_256.s
var assemblyVhaddpd256 string

//go:embed stub_vhaddpd_256.go
var stubVhaddpd256 string

type VHADDPD256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHADDPD256() *VHADDPD256 {
	return &VHADDPD256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 64),
		vals2: number.NewNamedFloatParameter("vals2", 256, 64),
		ret:   number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VHADDPD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHADDPD256) Output() *number.Parameter {
	return v.ret
}

func (v *VHADDPD256) Name() string {
	return "VHADDPD (256 bit) "
}

func (v *VHADDPD256) Description() string {
	return "Horizontal add of packed double-precision floats. " +
		"Computes pairwise sums within each source (a0+a1) and (b0+b1), writes results to ret."
}

func (v *VHADDPD256) Stub() string {
	return stubVhaddpd256
}

func (v *VHADDPD256) Assembly() string {
	return assemblyVhaddpd256
}

func (v *VHADDPD256) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals1.FlatData()))
	vals2 := [4]float64{}
	copy(vals2[:], number.ToFloat64Slice(v.vals2.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vhaddpd256(&vals1, &vals2, &ret)

	log.Printf("VHADDPD vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHADDPD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
