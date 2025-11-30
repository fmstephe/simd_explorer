package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_ord.s
var assemblyCmpss128Ord string

//go:embed stub_cmpss_128_ord.go
var stubCmpss128Ord string

type CMPSS128ORD struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128ORD() *CMPSS128ORD {
	return &CMPSS128ORD{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128ORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128ORD) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128ORD) Name() string {
	return "CMPSS (128 bit) ord"
}

func (v *CMPSS128ORD) Description() string {
	return "Compare scalar single-precision (lane 0) for ordered (neither operand is NaN); result mask in lane 0."
}

func (v *CMPSS128ORD) Stub() string {
	return stubCmpss128Ord
}

func (v *CMPSS128ORD) Assembly() string {
	return assemblyCmpss128Ord
}

func (v *CMPSS128ORD) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Ord(&vals1, &vals2, &ret)

	log.Printf("CMPSS128ORD input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPSS128ORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
