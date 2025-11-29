package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_le.s
var assemblyCmpss128Le string

//go:embed stub_cmpss_128_le.go
var stubCmpss128Le string

type CMPSS128LE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128LE() *CMPSS128LE {
	return &CMPSS128LE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128LE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128LE) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128LE) Name() string {
	return "CMPSS (128 bit) le"
}

func (v *CMPSS128LE) Description() string {
	return "Compare scalar single-precision (lane 0) for less-than-or-equal; result mask in lane 0, upper lanes pass through."
}

func (v *CMPSS128LE) Stub() string {
	return stubCmpss128Le
}

func (v *CMPSS128LE) Assembly() string {
	return assemblyCmpss128Le
}

func (v *CMPSS128LE) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Le(&vals1, &vals2, &ret)

	log.Printf("CMPSS128LE input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPSS128LE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
