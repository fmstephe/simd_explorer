package andnps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_andnps_128.s
var assemblyAndnps128 string

//go:embed stub_andnps_128.go
var stubAndnps128 string

type ANDNPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewANDNPS128() *ANDNPS128 {
	return &ANDNPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *ANDNPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *ANDNPS128) Output() *number.Parameter {
	return v.ret
}

func (v *ANDNPS128) Name() string {
	return "ANDNPS (128 bit)"
}

func (v *ANDNPS128) Description() string {
	return "Bitwise AND NOT of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *ANDNPS128) Stub() string {
	return stubAndnps128
}

func (v *ANDNPS128) Assembly() string {
	return assemblyAndnps128
}

func (v *ANDNPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	andnps128(&vals1, &vals2, &ret)

	log.Printf("ANDNPS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *ANDNPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
