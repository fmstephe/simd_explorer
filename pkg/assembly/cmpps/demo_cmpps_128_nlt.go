package cmpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpps_128_nlt.s
var assemblyCmpps128Nlt string

//go:embed stub_cmpps_128_nlt.go
var stubCmpps128Nlt string

type CMPPS128NLT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPPS128NLT() *CMPPS128NLT {
	return &CMPPS128NLT{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPPS128NLT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPPS128NLT) Output() *number.Parameter {
	return v.ret
}

func (v *CMPPS128NLT) Name() string {
	return "CMPPS (128 bit) nlt"
}

func (v *CMPPS128NLT) Description() string {
	return "Compare packed single-precision floats for not-less-than (per lane)."
}

func (v *CMPPS128NLT) Stub() string {
	return stubCmpps128Nlt
}

func (v *CMPPS128NLT) Assembly() string {
	return assemblyCmpps128Nlt
}

func (v *CMPPS128NLT) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpps128Nlt(&vals1, &vals2, &ret)

	log.Printf("CMPPS128NLT input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPPS128NLT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
