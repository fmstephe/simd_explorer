package andps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vandps_128.s
var assemblyVandps128 string

//go:embed stub_vandps_128.go
var stubVandps128 string

type VANDPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVANDPS128() *VANDPS128 {
	return &VANDPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VANDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VANDPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VANDPS128) Name() string {
	return "VANDPS (128 bit)"
}

func (v *VANDPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VANDPS128) Stub() string {
	return stubVandps128
}

func (v *VANDPS128) Assembly() string {
	return assemblyVandps128
}

func (v *VANDPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vandps128(&vals1, &vals2, &ret)

	log.Printf("VANDPS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VANDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
