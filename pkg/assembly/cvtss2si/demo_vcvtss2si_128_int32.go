package cvtss2si

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtss2si_128_int32.s
var assemblyVcvtss2si128Int32 string

//go:embed stub_vcvtss2si_128_int32.go
var stubVcvtss2si128Int32 string

type VCVTSS2SI128INT32 struct {
}

func (v *VCVTSS2SI128INT32) Inputs() []*number.Parameter {
	return []*number.Parameter{
		number.NewFloatParameter(128, 32),
	}
}

func (v *VCVTSS2SI128INT32) Output() *number.Parameter {
	return number.NewIntParameter(32, 32, 10)
}

func (v *VCVTSS2SI128INT32) Name() string {
	return "VCVTSS2SI (128 bit) int32"
}

func (v *VCVTSS2SI128INT32) Description() string {
	return "Convert scalar single-precision (lowest lane) to signed 32-bit integer."
}

func (v *VCVTSS2SI128INT32) Stub() string {
	return stubVcvtss2si128Int32
}

func (v *VCVTSS2SI128INT32) Assembly() string {
	return assemblyVcvtss2si128Int32
}

func (v *VCVTSS2SI128INT32) Run(inputs [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(inputs[0]))

	var ret int32

	vcvtss2si128Int32(&vals, &ret)

	log.Printf("VCVTSS2SI128INT32 input %v output %d", vals, ret)

	return number.Int32ToBytes(ret)
}

func (v *VCVTSS2SI128INT32) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
