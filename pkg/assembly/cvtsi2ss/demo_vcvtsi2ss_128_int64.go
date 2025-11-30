package cvtsi2ss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtsi2ss_128_int64.s
var assemblyVcvtsi2ss128Int64 string

//go:embed stub_vcvtsi2ss_128_int64.go
var stubVcvtsi2ss128Int64 string

type VCVTSI2SS128INT64 struct {
	vals *number.Parameter
	ival *number.Parameter
	ret  *number.Parameter
}

func NewVCVTSI2SS128INT64() *VCVTSI2SS128INT64 {
	return &VCVTSI2SS128INT64{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ival: number.NewNamedIntParameter("ival", 64, 64, 10),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VCVTSI2SS128INT64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals, // base vector (a0..a3)
		v.ival, // signed 64-bit integer
	}
}

func (v *VCVTSI2SS128INT64) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTSI2SS128INT64) Name() string {
	return "VCVTSI2SS (128 bit) int64->float32"
}

func (v *VCVTSI2SS128INT64) Description() string {
	return "Convert signed 64-bit integer to scalar single and insert into lowest lane; upper lanes preserved."
}

func (v *VCVTSI2SS128INT64) Stub() string {
	return stubVcvtsi2ss128Int64
}

func (v *VCVTSI2SS128INT64) Assembly() string {
	return assemblyVcvtsi2ss128Int64
}

func (v *VCVTSI2SS128INT64) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ival := number.ToInt64(v.ival.FlatData())

	ret := [4]float32{}

	vcvtsi2ss128int64(&vals, &ival, &ret)

	log.Printf("VCVTSI2SS128 input base %v int %d output %v", vals, ival, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VCVTSI2SS128INT64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
