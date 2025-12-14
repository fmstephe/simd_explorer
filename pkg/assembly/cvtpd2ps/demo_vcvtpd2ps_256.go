package cvtpd2ps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtpd2ps_256.s
var assemblyVcvtpd2ps256 string

//go:embed stub_vcvtpd2ps_256.go
var stubVcvtpd2ps256 string

type VCVTPD2PS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPD2PS256() *VCVTPD2PS256 {
	return &VCVTPD2PS256{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VCVTPD2PS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPD2PS256) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPD2PS256) Name() string {
	return "VCVTPD2PS (256 bit) "
}

func (v *VCVTPD2PS256) Description() string {
	return "Convert packed double-precision to single-precision. " +
		"Converts 4 float64 values from vals (YMM source) to 4 float32 results in ret (XMM destination)."
}

func (v *VCVTPD2PS256) Stub() string {
	return stubVcvtpd2ps256
}

func (v *VCVTPD2PS256) Assembly() string {
	return assemblyVcvtpd2ps256
}

func (v *VCVTPD2PS256) Run() {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))
	ret := [4]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vcvtpd2ps256(&vals, &ret)

	log.Printf("VCVTPD2PS vals %v ret %v", vals, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPD2PS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
