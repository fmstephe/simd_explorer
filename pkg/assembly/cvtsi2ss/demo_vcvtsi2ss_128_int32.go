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
}

func (v *VCVTSI2SS128INT32) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),  // base vector (a0..a3)
		number.NewIntParameter(32, 32, 10), // signed 32-bit integer
	}
}

func (v *VCVTSI2SS128INT32) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
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

func (v *VCVTSI2SS128INT32) Run(inputs [][]byte) (output []byte) {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(inputs[0]))
	intScalar := number.ToInt32(inputs[1])

	ret := [4]float32{}

	vcvtsi2ss128int32(&base, &intScalar, &ret)

	log.Printf("VCVTSI2SS128 input base %v int %d output %v", base, intScalar, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCVTSI2SS128INT32) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
