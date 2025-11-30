package andnps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandnps_128.s
var assemblyVandnps128 string

//go:embed stub_vandnps_128.go
var stubVandnps128 string

type VANDNPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVANDNPS128() *VANDNPS128 {
	return &VANDNPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VANDNPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VANDNPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VANDNPS128) Name() string {
	return "VANDNPS (128 bit)"
}

func (v *VANDNPS128) Description() string {
	return "Bitwise AND NOT with VEX encoding; output shown as 32-bit hex lanes."
}

func (v *VANDNPS128) Stub() string {
	return stubVandnps128
}

func (v *VANDNPS128) Assembly() string {
	return assemblyVandnps128
}

func (v *VANDNPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vandnps128(&vals1, &vals2, &ret)

	log.Printf("VANDNPS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *VANDNPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
