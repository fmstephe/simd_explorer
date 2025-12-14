package cvtps2pd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtps2pd_128.s
var assemblyVcvtps2pd128 string

//go:embed stub_vcvtps2pd_128.go
var stubVcvtps2pd128 string

type VCVTPS2PD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPS2PD128() *VCVTPS2PD128 {
	return &VCVTPS2PD128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VCVTPS2PD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPS2PD128) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPS2PD128) Name() string {
	return "VCVTPS2PD (128 bit) "
}

func (v *VCVTPS2PD128) Description() string {
	return "Convert packed single-precision floats to packed double-precision. " +
		"Uses the lower 2 float32 lanes of vals and writes 2 float64 results to ret; the upper 2 lanes of vals are ignored."
}

func (v *VCVTPS2PD128) Stub() string {
	return stubVcvtps2pd128
}

func (v *VCVTPS2PD128) Assembly() string {
	return assemblyVcvtps2pd128
}

func (v *VCVTPS2PD128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vcvtps2pd128(&vals, &ret)

	log.Printf("VCVTPS2PD vals %v ret %v", vals, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPS2PD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
