package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_le.s
var assemblyCmpps128Le string

//go:embed stub_cmpps_128_le.go
var stubCmpps128Le string

type CMPPS128LE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128LE() *CMPPS128LE {
	return &CMPPS128LE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128LE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128LE) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128LE) Name() string {
	return "CMPPS (128 bit) le"
}

func (v *CMPPS128LE) Description() string {
	return "Compare packed single-precision floats for less-than-or-equal (per lane)."
}

func (v *CMPPS128LE) Stub() string {
	return stubCmpps128Le
}

func (v *CMPPS128LE) Assembly() string {
	return assemblyCmpps128Le
}

func (v *CMPPS128LE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Le(&vals1, &vals2, &ret)

	log.Printf("CMPPS128LE input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *CMPPS128LE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
