package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_nle.s
var assemblyCmpps128Nle string

//go:embed stub_cmpps_128_nle.go
var stubCmpps128Nle string

type CMPPS128NLE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128NLE() *CMPPS128NLE {
	return &CMPPS128NLE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128NLE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128NLE) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128NLE) Name() string {
	return "CMPPS (128 bit) nle"
}

func (v *CMPPS128NLE) Description() string {
	return "Compare packed single-precision floats for not-less-or-equal (per lane)."
}

func (v *CMPPS128NLE) Stub() string {
	return stubCmpps128Nle
}

func (v *CMPPS128NLE) Assembly() string {
	return assemblyCmpps128Nle
}

func (v *CMPPS128NLE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Nle(&vals1, &vals2, &ret)

	log.Printf("CMPPS128NLE input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPPS128NLE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
