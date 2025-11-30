package sqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsqrtps_256.s
var assemblyVsqrtps256 string

//go:embed stub_vsqrtps_256.go
var stubVsqrtps256 string

type VSQRTPS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVSQRTPS256() *VSQRTPS256 {
	return &VSQRTPS256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSQRTPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VSQRTPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VSQRTPS256) Name() string {
	return "VSQRTPS YMM (256 bit)"
}

func (v *VSQRTPS256) Description() string {
	return "AVX form: compute square root of packed single-precision floats in YMM, lane-wise."
}

func (v *VSQRTPS256) Stub() string {
	return stubVsqrtps256
}

func (v *VSQRTPS256) Assembly() string {
	return assemblyVsqrtps256
}

func (v *VSQRTPS256) Run() (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vsqrtps256(&vals, &ret)

	log.Printf("VSQRTPS256 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSQRTPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
