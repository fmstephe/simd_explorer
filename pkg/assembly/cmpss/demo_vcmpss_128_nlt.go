package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_nlt.s
var assemblyVcmpss128Nlt string

//go:embed stub_vcmpss_128_nlt.go
var stubVcmpss128Nlt string

type VCMPSS128NLT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCMPSS128NLT() *VCMPSS128NLT {
	return &VCMPSS128NLT{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VCMPSS128NLT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCMPSS128NLT) Output() *number.Parameter {
	return v.ret
}

func (v *VCMPSS128NLT) Name() string {
	return "VCMPSS (128 bit) nlt"
}

func (v *VCMPSS128NLT) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for not-less-than; result mask in lane 0, upper lanes pass through."
}

func (v *VCMPSS128NLT) Stub() string {
	return stubVcmpss128Nlt
}

func (v *VCMPSS128NLT) Assembly() string {
	return assemblyVcmpss128Nlt
}

func (v *VCMPSS128NLT) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vcmpss128Nlt(&vals1, &vals2, &ret)

	log.Printf("VCMPSS128NLT input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VCMPSS128NLT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
