package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_neq.s
var assemblyCmpss128Neq string

//go:embed stub_cmpss_128_neq.go
var stubCmpss128Neq string

type CMPSS128NEQ struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128NEQ() *CMPSS128NEQ {
	return &CMPSS128NEQ{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128NEQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128NEQ) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128NEQ) Name() string {
	return "CMPSS (128 bit) neq"
}

func (v *CMPSS128NEQ) Description() string {
	return "Compare scalar single-precision (lane 0) for inequality; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128NEQ) Stub() string {
	return stubCmpss128Neq
}

func (v *CMPSS128NEQ) Assembly() string {
	return assemblyCmpss128Neq
}

func (v *CMPSS128NEQ) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Neq(&vals1, &vals2, &ret)

	log.Printf("CMPSS128NEQ input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPSS128NEQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
