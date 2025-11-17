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
}

func (v *VCVTSI2SS128INT64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),  // base vector (a0..a3)
		number.NewIntParameter(64, 64, 10), // signed 64-bit integer
	}
}

func (v *VCVTSI2SS128INT64) Output() *number.Parameter {
	return number.NewFloatParameter(128, 32)
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

func (v *VCVTSI2SS128INT64) Run(inputs [][]byte) (output []byte) {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(inputs[0]))
	intScalar := number.ToInt64(inputs[1])

	ret := [4]float32{}

	vcvtsi2ss128int64(&base, &intScalar, &ret)

	log.Printf("VCVTSI2SS128 input base %v int %d output %v", base, intScalar, ret)

	return number.Float32SliceToBytes(ret[:])
}

func (v *VCVTSI2SS128INT64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
