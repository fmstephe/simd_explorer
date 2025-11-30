package cmpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcmpss_128_unord.s
var assemblyVcmpss128Unord string

//go:embed stub_vcmpss_128_unord.go
var stubVcmpss128Unord string

type VCMPSS128UNORD struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVCMPSS128UNORD() *VCMPSS128UNORD {
	return &VCMPSS128UNORD{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VCMPSS128UNORD) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VCMPSS128UNORD) Output() *number.Parameter {
	return v.ret
}

func (v *VCMPSS128UNORD) Name() string {
	return "VCMPSS (128 bit) unord"
}

func (v *VCMPSS128UNORD) Description() string {
	return "AVX form: compare scalar single-precision (lane 0) for unordered (either operand is NaN); result mask in lane 0."
}

func (v *VCMPSS128UNORD) Stub() string {
	return stubVcmpss128Unord
}

func (v *VCMPSS128UNORD) Assembly() string {
	return assemblyVcmpss128Unord
}

func (v *VCMPSS128UNORD) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vcmpss128Unord(&vals1, &vals2, &ret)

	log.Printf("VCMPSS128UNORD input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *VCMPSS128UNORD) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
