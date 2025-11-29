package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_nlt.s
var assemblyCmpss128Nlt string

//go:embed stub_cmpss_128_nlt.go
var stubCmpss128Nlt string

type CMPSS128NLT struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128NLT() *CMPSS128NLT {
	return &CMPSS128NLT{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128NLT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128NLT) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128NLT) Name() string {
	return "CMPSS (128 bit) nlt"
}

func (v *CMPSS128NLT) Description() string {
	return "Compare scalar single-precision (lane 0) for not-less-than; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128NLT) Stub() string {
	return stubCmpss128Nlt
}

func (v *CMPSS128NLT) Assembly() string {
	return assemblyCmpss128Nlt
}

func (v *CMPSS128NLT) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Nlt(&vals1, &vals2, &ret)

	log.Printf("CMPSS128NLT input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPSS128NLT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
