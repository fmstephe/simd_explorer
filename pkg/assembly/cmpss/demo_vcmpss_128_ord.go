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
}

func (v *VCMPSS128ORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCMPSS128ORD) Output() *number.Parameter {
	return number.NewUintParameter(128, 32, 16)
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

func (v *VCMPSS128ORD) Run(inputs [][]byte) (output []byte) {
	floats1 := [4]float32{}
	copy(floats1[:], number.ToFloat32Slice(inputs[0]))
	floats2 := [4]float32{}
	copy(floats2[:], number.ToFloat32Slice(inputs[1]))

	ret := [4]float32{}

	vcmpss128Ord(&floats1, &floats2, &ret)

	log.Printf("VCMPSS128ORD input %v %v output %v", floats1, floats2, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCMPSS128ORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
