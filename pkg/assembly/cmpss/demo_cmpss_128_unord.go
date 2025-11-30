package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_cmpss_128_unord.s
var assemblyCmpss128Unord string

//go:embed stub_cmpss_128_unord.go
var stubCmpss128Unord string

type CMPSS128UNORD struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewCMPSS128UNORD() *CMPSS128UNORD {
	return &CMPSS128UNORD{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *CMPSS128UNORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *CMPSS128UNORD) Output() *number.Parameter {
	return v.ret
}

func (v *CMPSS128UNORD) Name() string {
	return "CMPSS (128 bit) unord"
}

func (v *CMPSS128UNORD) Description() string {
	return "Compare scalar single-precision (lane 0) for unordered (either operand is NaN); result mask in lane 0."
}

func (v *CMPSS128UNORD) Stub() string {
	return stubCmpss128Unord
}

func (v *CMPSS128UNORD) Assembly() string {
	return assemblyCmpss128Unord
}

func (v *CMPSS128UNORD) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	cmpss128Unord(&vals1, &vals2, &ret)

	log.Printf("CMPSS128UNORD input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *CMPSS128UNORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
