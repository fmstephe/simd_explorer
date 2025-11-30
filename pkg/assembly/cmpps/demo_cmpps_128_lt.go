package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_lt.s
var assemblyCmpps128Lt string

//go:embed stub_cmpps_128_lt.go
var stubCmpps128Lt string

type CMPPS128LT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128LT() *CMPPS128LT {
	return &CMPPS128LT{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128LT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128LT) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128LT) Name() string {
	return "CMPPS (128 bit) lt"
}

func (v *CMPPS128LT) Description() string {
	return "Compare packed single-precision floats for less-than (per lane)."
}

func (v *CMPPS128LT) Stub() string {
	return stubCmpps128Lt
}

func (v *CMPPS128LT) Assembly() string {
	return assemblyCmpps128Lt
}

func (v *CMPPS128LT) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Lt(&vals1, &vals2, &ret)

	log.Printf("CMPPS128LT input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPPS128LT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
