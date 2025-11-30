package cvtsi2ss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtsi2ss_128_int32.s
var assemblyVcvtsi2ss128Int32 string

//go:embed stub_vcvtsi2ss_128_int32.go
var stubVcvtsi2ss128Int32 string

type VCVTSI2SS128INT32 struct {
	vals *number.Parameter
	ival *number.Parameter
	ret  *number.Parameter
}

func NewVCVTSI2SS128INT32() *VCVTSI2SS128INT32 {
	return &VCVTSI2SS128INT32{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ival: number.NewNamedIntParameter("ival", 32, 32, 10),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VCVTSI2SS128INT32) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals, // base vector (a0..a3)
		v.ival, // signed 32-bit integer
	}
}

func (v *VCVTSI2SS128INT32) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTSI2SS128INT32) Name() string {
	return "VCVTSI2SS (128 bit) int32->float32"
}

func (v *VCVTSI2SS128INT32) Description() string {
	return "Convert signed 32-bit integer to scalar single and insert into lowest lane; upper lanes preserved."
}

func (v *VCVTSI2SS128INT32) Stub() string {
	return stubVcvtsi2ss128Int32
}

func (v *VCVTSI2SS128INT32) Assembly() string {
	return assemblyVcvtsi2ss128Int32
}

func (v *VCVTSI2SS128INT32) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))
	ival := number.ToInt32(v.ival.FlatData())

	ret := [4]float32{}

	vcvtsi2ss128int32(&vals, &ival, &ret)

	log.Printf("VCVTSI2SS128 input base %v int %d output %v", vals, ival, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VCVTSI2SS128INT32) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
