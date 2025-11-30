package rsqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrsqrtps_128.s
var assemblyVrsqrtps128 string

//go:embed stub_vrsqrtps_128.go
var stubVrsqrtps128 string

type VRSQRTPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVRSQRTPS128() *VRSQRTPS128 {
	return &VRSQRTPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VRSQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VRSQRTPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VRSQRTPS128) Name() string {
	return "VRSQRTPS XMM (128 bit)"
}

func (v *VRSQRTPS128) Description() string {
	return "AVX form: reciprocal square root estimate of packed single-precision floats in XMM, lane-wise."
}

func (v *VRSQRTPS128) Stub() string {
	return stubVrsqrtps128
}

func (v *VRSQRTPS128) Assembly() string {
	return assemblyVrsqrtps128
}

func (v *VRSQRTPS128) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vrsqrtps128(&vals, &ret)

	log.Printf("VRSQRTPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VRSQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
