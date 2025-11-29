package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_ord.s
var assemblyVcmpss128Ord string

//go:embed stub_vcmpss_128_ord.go
var stubVcmpss128Ord string

type VCMPSS128ORD struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCMPSS128ORD() *VCMPSS128ORD {
	return &VCMPSS128ORD{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VCMPSS128ORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCMPSS128ORD) Output() *number.Parameter {
	return v.ret
}

func (v *VCMPSS128ORD) Name() string {
	return "VCMPSS (128 bit) ord"
}

func (v *VCMPSS128ORD) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for ordered (neither operand is NaN); result mask in lane 0."
}

func (v *VCMPSS128ORD) Stub() string {
	return stubVcmpss128Ord
}

func (v *VCMPSS128ORD) Assembly() string {
	return assemblyVcmpss128Ord
}

func (v *VCMPSS128ORD) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vcmpss128Ord(&vals1, &vals2, &ret)

	log.Printf("VCMPSS128ORD input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VCMPSS128ORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
