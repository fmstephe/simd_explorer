package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_lt.s
var assemblyCmpss128Lt string

//go:embed stub_cmpss_128_lt.go
var stubCmpss128Lt string

type CMPSS128LT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128LT() *CMPSS128LT {
	return &CMPSS128LT{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128LT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128LT) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128LT) Name() string {
	return "CMPSS (128 bit) lt"
}

func (v *CMPSS128LT) Description() string {
	return "Compare scalar single-precision (lane 0) for less-than; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128LT) Stub() string {
	return stubCmpss128Lt
}

func (v *CMPSS128LT) Assembly() string {
	return assemblyCmpss128Lt
}

func (v *CMPSS128LT) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Lt(&vals1, &vals2, &ret)

	log.Printf("CMPSS128LT input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPSS128LT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
