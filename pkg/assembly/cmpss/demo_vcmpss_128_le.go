package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_le.s
var assemblyVcmpss128Le string

//go:embed stub_vcmpss_128_le.go
var stubVcmpss128Le string

type VCMPSS128LE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCMPSS128LE() *VCMPSS128LE {
	return &VCMPSS128LE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VCMPSS128LE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCMPSS128LE) Output() *number.Parameter {
	return v.ret
}

func (v *VCMPSS128LE) Name() string {
	return "VCMPSS (128 bit) le"
}

func (v *VCMPSS128LE) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for less-than-or-equal; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128LE) Stub() string {
	return stubVcmpss128Le
}

func (v *VCMPSS128LE) Assembly() string {
	return assemblyVcmpss128Le
}

func (v *VCMPSS128LE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vcmpss128Le(&vals1, &vals2, &ret)

	log.Printf("VCMPSS128LE input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *VCMPSS128LE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
