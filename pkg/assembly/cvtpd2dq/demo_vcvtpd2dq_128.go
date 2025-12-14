package cvtpd2dq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtpd2dq_128.s
var assemblyVcvtpd2dq128 string

//go:embed stub_vcvtpd2dq_128.go
var stubVcvtpd2dq128 string

type VCVTPD2DQ128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPD2DQ128() *VCVTPD2DQ128 {
	return &VCVTPD2DQ128{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VCVTPD2DQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPD2DQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPD2DQ128) Name() string {
	return "VCVTPD2DQ (128 bit) "
}

func (v *VCVTPD2DQ128) Description() string {
	return "Convert packed double-precision floats to packed signed 32-bit integers. " +
		"Uses MXCSR rounding mode (default: round-to-nearest). " +
		"NaN or out-of-range inputs produce 0x80000000 for that element."
}

func (v *VCVTPD2DQ128) Stub() string {
	return stubVcvtpd2dq128
}

func (v *VCVTPD2DQ128) Assembly() string {
	return assemblyVcvtpd2dq128
}

func (v *VCVTPD2DQ128) Run() {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vcvtpd2dq128(&vals, &ret)

	log.Printf("VCVTPD2DQ vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPD2DQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
