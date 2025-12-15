package vaddsubps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddsubps_256.s
var assemblyVaddsubps256 string

//go:embed stub_vaddsubps_256.go
var stubVaddsubps256 string

type VADDSUBPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDSUBPS256() *VADDSUBPS256 {
	return &VADDSUBPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VADDSUBPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDSUBPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VADDSUBPS256) Name() string {
	return "VADDSUBPS (256 bit) "
}

func (v *VADDSUBPS256) Description() string {
	return "Add-subtract packed single-precision floats. " +
		"Even elements are added and odd elements subtracted within each 128-bit lane; results stored in ret."
}

func (v *VADDSUBPS256) Stub() string {
	return stubVaddsubps256
}

func (v *VADDSUBPS256) Assembly() string {
	return assemblyVaddsubps256
}

func (v *VADDSUBPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [8]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vaddsubps256(&vals1, &vals2, &ret)

	log.Printf("VADDSUBPS vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VADDSUBPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
