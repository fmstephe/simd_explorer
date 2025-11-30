package rsqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrsqrtps_256.s
var assemblyVrsqrtps256 string

//go:embed stub_vrsqrtps_256.go
var stubVrsqrtps256 string

type VRSQRTPS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVRSQRTPS256() *VRSQRTPS256 {
	return &VRSQRTPS256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VRSQRTPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VRSQRTPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VRSQRTPS256) Name() string {
	return "VRSQRTPS YMM (256 bit)"
}

func (v *VRSQRTPS256) Description() string {
	return "AVX form: reciprocal square root estimate of packed single-precision floats in YMM, lane-wise."
}

func (v *VRSQRTPS256) Stub() string {
	return stubVrsqrtps256
}

func (v *VRSQRTPS256) Assembly() string {
	return assemblyVrsqrtps256
}

func (v *VRSQRTPS256) Run() {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vrsqrtps256(&vals, &ret)

	log.Printf("VRSQRTPS256 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VRSQRTPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
