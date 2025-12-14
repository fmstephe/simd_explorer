package cvtdq2ps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtdq2ps_128.s
var assemblyVcvtdq2ps128 string

//go:embed stub_vcvtdq2ps_128.go
var stubVcvtdq2ps128 string

type VCVTDQ2PS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTDQ2PS128() *VCVTDQ2PS128 {
	return &VCVTDQ2PS128{
		vals: number.NewNamedIntParameter("vals", 128, 32, 10),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VCVTDQ2PS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTDQ2PS128) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTDQ2PS128) Name() string {
	return "VCVTDQ2PS (128 bit) "
}

func (v *VCVTDQ2PS128) Description() string {
	return "Convert packed signed 32-bit integers to packed single-precision floats. " +
		"Converts 4 int32 values from vals to 4 float32 results in ret."
}

func (v *VCVTDQ2PS128) Stub() string {
	return stubVcvtdq2ps128
}

func (v *VCVTDQ2PS128) Assembly() string {
	return assemblyVcvtdq2ps128
}

func (v *VCVTDQ2PS128) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [4]float32{}
	copy(ret[:], number.ToFloat32Slice(v.ret.FlatData()))

	vcvtdq2ps128(&vals, &ret)

	log.Printf("VCVTDQ2PS vals %v ret %v", vals, ret)

	retBytes := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTDQ2PS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
