package cvtdq2pd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtdq2pd_256.s
var assemblyVcvtdq2pd256 string

//go:embed stub_vcvtdq2pd_256.go
var stubVcvtdq2pd256 string

type VCVTDQ2PD256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTDQ2PD256() *VCVTDQ2PD256 {
	return &VCVTDQ2PD256{
		vals: number.NewNamedIntParameter("vals", 128, 32, 10),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VCVTDQ2PD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTDQ2PD256) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTDQ2PD256) Name() string {
	return "VCVTDQ2PD (256 bit) "
}

func (v *VCVTDQ2PD256) Description() string {
	return "Convert packed signed 32-bit integers to packed double-precision floats. " +
		"Converts 4 int32 lanes from vals to 4 float64 results in ret."
}

func (v *VCVTDQ2PD256) Stub() string {
	return stubVcvtdq2pd256
}

func (v *VCVTDQ2PD256) Assembly() string {
	return assemblyVcvtdq2pd256
}

func (v *VCVTDQ2PD256) Run() {
	vals := [4]int32{}
	copy(vals[:], number.ToInt32Slice(v.vals.FlatData()))
	ret := [4]float64{}
	copy(ret[:], number.ToFloat64Slice(v.ret.FlatData()))

	vcvtdq2pd256(&vals, &ret)

	log.Printf("VCVTDQ2PD vals %v ret %v", vals, ret)

	retBytes := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VCVTDQ2PD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
