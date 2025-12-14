package cvtdq2pd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtdq2pd_128.s
var assemblyVcvtdq2pd128 string

//go:embed stub_vcvtdq2pd_128.go
var stubVcvtdq2pd128 string

type VCVTDQ2PD128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTDQ2PD128() *VCVTDQ2PD128 {
	return &VCVTDQ2PD128{
		vals: number.NewNamedIntParameter("vals", 128, 32, 10),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VCVTDQ2PD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTDQ2PD128) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTDQ2PD128) Name() string {
	return "VCVTDQ2PD (128 bit) "
}

func (v *VCVTDQ2PD128) Description() string {
	return "Convert packed signed 32-bit integers to packed double-precision floats. " +
		"Converts the lower 2 int32 lanes from vals to 2 float64 results in ret; the upper 2 int32 lanes are ignored."
}

func (v *VCVTDQ2PD128) Stub() string {
	return stubVcvtdq2pd128
}

func (v *VCVTDQ2PD128) Assembly() string {
	return assemblyVcvtdq2pd128
}

func (v *VCVTDQ2PD128) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [2]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vcvtdq2pd128(&vals, &ret)

	log.Printf("VCVTDQ2PD vals %v ret %v", vals, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTDQ2PD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
