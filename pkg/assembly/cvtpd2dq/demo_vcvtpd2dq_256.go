package cvtpd2dq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtpd2dq_256.s
var assemblyVcvtpd2dq256 string

//go:embed stub_vcvtpd2dq_256.go
var stubVcvtpd2dq256 string

type VCVTPD2DQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPD2DQ256() *VCVTPD2DQ256 {
	return &VCVTPD2DQ256{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedIntParameter("ret", 128, 32, 10),
	}
}

func (v *VCVTPD2DQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPD2DQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPD2DQ256) Name() string {
	return "VCVTPD2DQ (256 bit) "
}

func (v *VCVTPD2DQ256) Description() string {
	return "Convert packed double-precision floats to packed signed 32-bit integers. " +
		"Converts 4 lanes using the MXCSR rounding mode (default: round-to-nearest). " +
		"NaN or out-of-range inputs produce 0x80000000 for that element."
}

func (v *VCVTPD2DQ256) Stub() string {
	return stubVcvtpd2dq256
}

func (v *VCVTPD2DQ256) Assembly() string {
	return assemblyVcvtpd2dq256
}

func (v *VCVTPD2DQ256) Run() {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	ret := [4]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vcvtpd2dq256(&vals, &ret)

	log.Printf("VCVTPD2DQ vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPD2DQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
