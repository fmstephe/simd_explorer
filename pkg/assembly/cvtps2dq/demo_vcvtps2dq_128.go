package cvtps2dq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtps2dq_128.s
var assemblyVcvtps2dq128 string

//go:embed stub_vcvtps2dq_128.go
var stubVcvtps2dq128 string

type VCVTPS2DQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPS2DQ128() *VCVTPS2DQ128 {
	return &VCVTPS2DQ128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VCVTPS2DQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPS2DQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPS2DQ128) Name() string {
	return "VCVTPS2DQ (128 bit) "
}

func (v *VCVTPS2DQ128) Description() string {
	return "Convert packed single-precision floats to packed signed 32-bit integers. " +
		"Uses MXCSR rounding mode (default: round-to-nearest). " +
		"NaN or out-of-range inputs produce 0x80000000 for that element."
}

func (v *VCVTPS2DQ128) Stub() string {
	return stubVcvtps2dq128
}

func (v *VCVTPS2DQ128) Assembly() string {
	return assemblyVcvtps2dq128
}

func (v *VCVTPS2DQ128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vcvtps2dq128(&vals, &ret)

	log.Printf("VCVTPS2DQ vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPS2DQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
