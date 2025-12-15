package vhsubps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vhsubps_256.s
var assemblyVhsubps256 string

//go:embed stub_vhsubps_256.go
var stubVhsubps256 string

type VHSUBPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVHSUBPS256() *VHSUBPS256 {
	return &VHSUBPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VHSUBPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VHSUBPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VHSUBPS256) Name() string {
	return "VHSUBPS (256 bit) "
}

func (v *VHSUBPS256) Description() string {
	return "Horizontal subtract of packed single-precision floats. " +
		"Computes pairwise differences within each 128-bit lane of vals1 and vals2; results stored in ret."
}

func (v *VHSUBPS256) Stub() string {
	return stubVhsubps256
}

func (v *VHSUBPS256) Assembly() string {
	return assemblyVhsubps256
}

func (v *VHSUBPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [8]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vhsubps256(&vals1, &vals2, &ret)

	log.Printf("VHSUBPS vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VHSUBPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
