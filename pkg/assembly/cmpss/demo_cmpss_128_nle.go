package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_nle.s
var assemblyCmpss128Nle string

//go:embed stub_cmpss_128_nle.go
var stubCmpss128Nle string

type CMPSS128NLE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128NLE() *CMPSS128NLE {
	return &CMPSS128NLE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128NLE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128NLE) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128NLE) Name() string {
	return "CMPSS (128 bit) nle"
}

func (v *CMPSS128NLE) Description() string {
	return "Compare scalar single-precision (lane 0) for not-less-or-equal; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128NLE) Stub() string {
	return stubCmpss128Nle
}

func (v *CMPSS128NLE) Assembly() string {
	return assemblyCmpss128Nle
}

func (v *CMPSS128NLE) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Nle(&vals1, &vals2, &ret)

	log.Printf("CMPSS128NLE input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPSS128NLE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
