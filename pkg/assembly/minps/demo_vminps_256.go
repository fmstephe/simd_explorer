package minps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vminps_256.s
var assemblyVminps256 string

//go:embed stub_vminps_256.go
var stubVminps256 string

type VMINPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVMINPS256() *VMINPS256 {
	return &VMINPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VMINPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VMINPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VMINPS256) Name() string {
	return "VMINPS YMM (256 bit)"
}

func (v *VMINPS256) Description() string {
	return "AVX form: compute element-wise minimum of packed single-precision floats in YMM."
}

func (v *VMINPS256) Stub() string {
	return stubVminps256
}

func (v *VMINPS256) Assembly() string {
	return assemblyVminps256
}

func (v *VMINPS256) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vminps256(&vals1, &vals2, &ret)

	log.Printf("VMINPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VMINPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
