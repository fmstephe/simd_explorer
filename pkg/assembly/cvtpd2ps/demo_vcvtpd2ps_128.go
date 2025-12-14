package cvtpd2ps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtpd2ps_128.s
var assemblyVcvtpd2ps128 string

//go:embed stub_vcvtpd2ps_128.go
var stubVcvtpd2ps128 string

type VCVTPD2PS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPD2PS128() *VCVTPD2PS128 {
	return &VCVTPD2PS128{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VCVTPD2PS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPD2PS128) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPD2PS128) Name() string {
	return "VCVTPD2PS (128 bit) "
}

func (v *VCVTPD2PS128) Description() string {
	return "Convert packed double-precision to single-precision. " +
		"Converts 2 float64 values from vals to 2 float32 results in ret; the upper 2 float32 lanes are zero."
}

func (v *VCVTPD2PS128) Stub() string {
	return stubVcvtpd2ps128
}

func (v *VCVTPD2PS128) Assembly() string {
	return assemblyVcvtpd2ps128
}

func (v *VCVTPD2PS128) Run() {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	ret := [4]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vcvtpd2ps128(&vals, &ret)

	log.Printf("VCVTPD2PS vals %v ret %v", vals, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPD2PS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
