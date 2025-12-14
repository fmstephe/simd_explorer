package cvtdq2ps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtdq2ps_256.s
var assemblyVcvtdq2ps256 string

//go:embed stub_vcvtdq2ps_256.go
var stubVcvtdq2ps256 string

type VCVTDQ2PS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTDQ2PS256() *VCVTDQ2PS256 {
	return &VCVTDQ2PS256{
		vals: number.NewNamedIntParameter("vals", 256, 32, 10),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VCVTDQ2PS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTDQ2PS256) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTDQ2PS256) Name() string {
	return "VCVTDQ2PS (256 bit) "
}

func (v *VCVTDQ2PS256) Description() string {
	return "Convert packed signed 32-bit integers to packed single-precision floats. " +
		"Converts 8 int32 values from vals to 8 float32 results in ret."
}

func (v *VCVTDQ2PS256) Stub() string {
	return stubVcvtdq2ps256
}

func (v *VCVTDQ2PS256) Assembly() string {
	return assemblyVcvtdq2ps256
}

func (v *VCVTDQ2PS256) Run() {
	vals := [8]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [8]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vcvtdq2ps256(&vals, &ret)

	log.Printf("VCVTDQ2PS vals %v ret %v", vals, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTDQ2PS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
