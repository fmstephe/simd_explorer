package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_nle.s
var assemblyVcmpss128Nle string

//go:embed stub_vcmpss_128_nle.go
var stubVcmpss128Nle string

type VCMPSS128NLE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCMPSS128NLE() *VCMPSS128NLE {
	return &VCMPSS128NLE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VCMPSS128NLE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCMPSS128NLE) Output() *number.Parameter {
	return v.ret
}

func (v *VCMPSS128NLE) Name() string {
	return "VCMPSS (128 bit) nle"
}

func (v *VCMPSS128NLE) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for not-less-or-equal; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128NLE) Stub() string {
	return stubVcmpss128Nle
}

func (v *VCMPSS128NLE) Assembly() string {
	return assemblyVcmpss128Nle
}

func (v *VCMPSS128NLE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vcmpss128Nle(&vals1, &vals2, &ret)

	log.Printf("VCMPSS128NLE input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *VCMPSS128NLE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
