package sqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_sqrtps_128.s
var assemblySqrtps128 string

//go:embed stub_sqrtps_128.go
var stubSqrtps128 string

type SQRTPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewSQRTPS128() *SQRTPS128 {
	return &SQRTPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *SQRTPS128) Output() *number.Parameter {
	return v.ret
}

func (v *SQRTPS128) Name() string {
	return "SQRTPS XMM (128 bit)"
}

func (v *SQRTPS128) Description() string {
	return "Compute square root of packed single-precision floats in XMM, lane-wise."
}

func (v *SQRTPS128) Stub() string {
	return stubSqrtps128
}

func (v *SQRTPS128) Assembly() string {
	return assemblySqrtps128
}

func (v *SQRTPS128) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	sqrtps128(&vals, &ret)

	log.Printf("SQRTPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
