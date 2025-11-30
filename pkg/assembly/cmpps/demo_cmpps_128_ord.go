package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_ord.s
var assemblyCmpps128Ord string

//go:embed stub_cmpps_128_ord.go
var stubCmpps128Ord string

type CMPPS128ORD struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128ORD() *CMPPS128ORD {
	return &CMPPS128ORD{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128ORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128ORD) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128ORD) Name() string {
	return "CMPPS (128 bit) ord"
}

func (v *CMPPS128ORD) Description() string {
	return "Compare packed single-precision floats for ordered (per lane)."
}

func (v *CMPPS128ORD) Stub() string {
	return stubCmpps128Ord
}

func (v *CMPPS128ORD) Assembly() string {
	return assemblyCmpps128Ord
}

func (v *CMPPS128ORD) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Ord(&vals1, &vals2, &ret)

	log.Printf("CMPPS128ORD input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPPS128ORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
