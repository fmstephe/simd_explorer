package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_eq.s
var assemblyCmpps128Eq string

//go:embed stub_cmpps_128_eq.go
var stubCmpps128Eq string

type CMPPS128EQ struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128EQ() *CMPPS128EQ {
	return &CMPPS128EQ{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128EQ) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128EQ) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128EQ) Name() string {
	return "CMPPS (128 bit) eq"
}

func (v *CMPPS128EQ) Description() string {
	return "Compare packed single-precision floats for equality (per lane)."
}

func (v *CMPPS128EQ) Stub() string {
	return stubCmpps128Eq
}

func (v *CMPPS128EQ) Assembly() string {
	return assemblyCmpps128Eq
}

func (v *CMPPS128EQ) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Eq(&vals1, &vals2, &ret)

	log.Printf("CMPPS128EQ input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPPS128EQ) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
