package minps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vminps_128.s
var assemblyVminps128 string

//go:embed stub_vminps_128.go
var stubVminps128 string

type VMINPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMINPS128() *VMINPS128 {
	return &VMINPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VMINPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMINPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VMINPS128) Name() string {
	return "VMINPS XMM (128 bit)"
}

func (v *VMINPS128) Description() string {
	return "AVX form: compute element-wise minimum of packed single-precision floats in XMM."
}

func (v *VMINPS128) Stub() string {
	return stubVminps128
}

func (v *VMINPS128) Assembly() string {
	return assemblyVminps128
}

func (v *VMINPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vminps128(&vals1, &vals2, &ret)

	log.Printf("VMINPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMINPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
