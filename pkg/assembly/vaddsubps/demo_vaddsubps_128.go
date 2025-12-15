package vaddsubps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddsubps_128.s
var assemblyVaddsubps128 string

//go:embed stub_vaddsubps_128.go
var stubVaddsubps128 string

type VADDSUBPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDSUBPS128() *VADDSUBPS128 {
	return &VADDSUBPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VADDSUBPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDSUBPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VADDSUBPS128) Name() string {
	return "VADDSUBPS (128 bit) "
}

func (v *VADDSUBPS128) Description() string {
	return "Add-subtract packed single-precision floats. " +
		"Even elements are added (a0+b0, a2+b2), odd elements are subtracted (a1-b1, a3-b3) and stored in ret."
}

func (v *VADDSUBPS128) Stub() string {
	return stubVaddsubps128
}

func (v *VADDSUBPS128) Assembly() string {
	return assemblyVaddsubps128
}

func (v *VADDSUBPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))
	ret := [4]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vaddsubps128(&vals1, &vals2, &ret)

	log.Printf("VADDSUBPS vals1 %v vals2 %v ret %v", vals1, vals2, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VADDSUBPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
