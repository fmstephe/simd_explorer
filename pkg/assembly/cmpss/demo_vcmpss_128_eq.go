package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_eq.s
var assemblyVcmpss128Eq string

//go:embed stub_vcmpss_128_eq.go
var stubVcmpss128Eq string

type VCMPSS128EQ struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCMPSS128EQ() *VCMPSS128EQ {
	return &VCMPSS128EQ{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VCMPSS128EQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCMPSS128EQ) Output() *number.Parameter {
	return v.ret
}

func (v *VCMPSS128EQ) Name() string {
	return "VCMPSS (128 bit) eq"
}

func (v *VCMPSS128EQ) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for equality; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128EQ) Stub() string {
	return stubVcmpss128Eq
}

func (v *VCMPSS128EQ) Assembly() string {
	return assemblyVcmpss128Eq
}

func (v *VCMPSS128EQ) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vcmpss128Eq(&vals1, &vals2, &ret)

	log.Printf("VCMPSS128EQ input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *VCMPSS128EQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
