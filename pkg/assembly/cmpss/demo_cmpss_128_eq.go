package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_eq.s
var assemblyCmpss128Eq string

//go:embed stub_cmpss_128_eq.go
var stubCmpss128Eq string

type CMPSS128EQ struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128EQ() *CMPSS128EQ {
	return &CMPSS128EQ{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128EQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128EQ) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128EQ) Name() string {
	return "CMPSS (128 bit) eq"
}

func (v *CMPSS128EQ) Description() string {
	return "Compare scalar single-precision (lane 0) for equality; writes 0xFFFFFFFF or 0x0 in lane 0, upper lanes pass through."
}

func (v *CMPSS128EQ) Stub() string {
	return stubCmpss128Eq
}

func (v *CMPSS128EQ) Assembly() string {
	return assemblyCmpss128Eq
}

func (v *CMPSS128EQ) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Eq(&vals1, &vals2, &ret)

	log.Printf("CMPSS128EQ input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPSS128EQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
