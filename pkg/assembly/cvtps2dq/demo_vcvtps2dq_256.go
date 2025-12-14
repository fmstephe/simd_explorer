package cvtps2dq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtps2dq_256.s
var assemblyVcvtps2dq256 string

//go:embed stub_vcvtps2dq_256.go
var stubVcvtps2dq256 string

type VCVTPS2DQ256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPS2DQ256() *VCVTPS2DQ256 {
	return &VCVTPS2DQ256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedIntParameter("ret", 256, 32, 10),
	}
}

func (v *VCVTPS2DQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPS2DQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPS2DQ256) Name() string {
	return "VCVTPS2DQ (256 bit) "
}

func (v *VCVTPS2DQ256) Description() string {
	return "Convert packed single-precision floats to packed signed 32-bit integers. " +
		"Converts 8 lanes using the MXCSR rounding mode (default: round-to-nearest). " +
		"NaN or out-of-range inputs produce 0x80000000 for that element."
}

func (v *VCVTPS2DQ256) Stub() string {
	return stubVcvtps2dq256
}

func (v *VCVTPS2DQ256) Assembly() string {
	return assemblyVcvtps2dq256
}

func (v *VCVTPS2DQ256) Run() {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [8]int32{}
	copy(ret[:], number.ToInt32Slice(v.ret.FlatData()))

	vcvtps2dq256(&vals, &ret)

	log.Printf("VCVTPS2DQ vals %v ret %v", vals, ret)

	retBytes := number.Int32SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPS2DQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
