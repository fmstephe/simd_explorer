package sqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsqrtps_128.s
var assemblyVsqrtps128 string

//go:embed stub_vsqrtps_128.go
var stubVsqrtps128 string

type VSQRTPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVSQRTPS128() *VSQRTPS128 {
	return &VSQRTPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VSQRTPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VSQRTPS128) Name() string {
	return "VSQRTPS XMM (128 bit)"
}

func (v *VSQRTPS128) Description() string {
	return "AVX form: compute square root of packed single-precision floats in XMM, lane-wise."
}

func (v *VSQRTPS128) Stub() string {
	return stubVsqrtps128
}

func (v *VSQRTPS128) Assembly() string {
	return assemblyVsqrtps128
}

func (v *VSQRTPS128) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vsqrtps128(&vals, &ret)

	log.Printf("VSQRTPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
