package addps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddps_256.s
var assemblyVaddps256 string

//go:embed stub_vaddps_256.go
var stubVaddps256 string

type VADDPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDPS256() *VADDPS256 {
	return &VADDPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VADDPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VADDPS256) Name() string {
	return "VADDPS YMM (256 bit)"
}

func (v *VADDPS256) Description() string {
	return "AVX form: add packed single-precision floats in YMM, lane-wise."
}

func (v *VADDPS256) Stub() string {
	return stubVaddps256
}

func (v *VADDPS256) Assembly() string {
	return assemblyVaddps256
}

func (v *VADDPS256) Run() (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vaddps256(&vals1, &vals2, &ret)

	log.Printf("VADDPS256 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VADDPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
