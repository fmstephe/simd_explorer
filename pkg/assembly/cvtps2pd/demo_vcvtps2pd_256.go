package cvtps2pd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtps2pd_256.s
var assemblyVcvtps2pd256 string

//go:embed stub_vcvtps2pd_256.go
var stubVcvtps2pd256 string

type VCVTPS2PD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTPS2PD256() *VCVTPS2PD256 {
	return &VCVTPS2PD256{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VCVTPS2PD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTPS2PD256) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTPS2PD256) Name() string {
	return "VCVTPS2PD (256 bit) "
}

func (v *VCVTPS2PD256) Description() string {
	return "Convert packed single-precision floats to packed double-precision. " +
		"Converts 4 float32 values from vals to 4 float64 values in ret (XMM source → YMM destination)."
}

func (v *VCVTPS2PD256) Stub() string {
	return stubVcvtps2pd256
}

func (v *VCVTPS2PD256) Assembly() string {
	return assemblyVcvtps2pd256
}

func (v *VCVTPS2PD256) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vcvtps2pd256(&vals, &ret)

	log.Printf("VCVTPS2PD vals %v ret %v", vals, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTPS2PD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
