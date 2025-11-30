package andps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_andps_128.s
var assemblyAndps128 string

//go:embed stub_andps_128.go
var stubAndps128 string

type ANDPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewANDPS128() *ANDPS128 {
	return &ANDPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *ANDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *ANDPS128) Output() *number.Parameter {
	return v.ret
}

func (v *ANDPS128) Name() string {
	return "ANDPS (128 bit)"
}

func (v *ANDPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *ANDPS128) Stub() string {
	return stubAndps128
}

func (v *ANDPS128) Assembly() string {
	return assemblyAndps128
}

func (v *ANDPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	andps128(&vals1, &vals2, &ret)

	log.Printf("ANDPS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)

}

func (v *ANDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
