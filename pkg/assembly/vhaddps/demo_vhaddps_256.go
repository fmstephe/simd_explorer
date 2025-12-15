package vhaddps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhaddps_256.s
var assemblyVhaddps256 string

//go:embed stub_vhaddps_256.go
var stubVhaddps256 string

type VHADDPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHADDPS256() *VHADDPS256 {
	return &VHADDPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VHADDPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHADDPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VHADDPS256) Name() string {
	return "VHADDPS (256 bit) "
}

func (v *VHADDPS256) Description() string {
	return "Horizontal add of packed single-precision floats. " +
		"Computes pairwise sums within each 128-bit lane of vals1 and vals2; results stored in ret."
}

func (v *VHADDPS256) Stub() string {
	return stubVhaddps256
}

func (v *VHADDPS256) Assembly() string {
	return assemblyVhaddps256
}

func (v *VHADDPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [8]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vhaddps256(&vals1, &vals2, &ret)

	log.Printf("VHADDPS vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHADDPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
