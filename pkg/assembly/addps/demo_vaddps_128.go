package addps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddps_128.s
var assemblyVaddps128 string

//go:embed stub_vaddps_128.go
var stubVaddps128 string

type VADDPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDPS128() *VADDPS128 {
	return &VADDPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VADDPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VADDPS128) Name() string {
	return "VADDPS XMM (128 bit)"
}

func (v *VADDPS128) Description() string {
	return "AVX form: add packed single-precision floats in XMM, lane-wise."
}

func (v *VADDPS128) Stub() string {
	return stubVaddps128
}

func (v *VADDPS128) Assembly() string {
	return assemblyVaddps128
}

func (v *VADDPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vaddps128(&vals1, &vals2, &ret)

	log.Printf("VADDPS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VADDPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
