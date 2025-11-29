package cvtss2si

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvtss2si_128_int64.s
var assemblyVcvtss2si128Int64 string

//go:embed stub_vcvtss2si_128_int64.go
var stubVcvtss2si128Int64 string

type VCVTSS2SI128INT64 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTSS2SI128INT64() *VCVTSS2SI128INT64 {
	return &VCVTSS2SI128INT64{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedIntParameter("ret", 64, 64, 10),
	}
}

func (v *VCVTSS2SI128INT64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTSS2SI128INT64) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTSS2SI128INT64) Name() string {
	return "VCVTSS2SI (128 bit) int64"
}

func (v *VCVTSS2SI128INT64) Description() string {
	return "Convert scalar single-precision (lowest lane) to signed 64-bit integer."
}

func (v *VCVTSS2SI128INT64) Stub() string {
	return stubVcvtss2si128Int64
}

func (v *VCVTSS2SI128INT64) Assembly() string {
	return assemblyVcvtss2si128Int64
}

func (v *VCVTSS2SI128INT64) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	var ret int64

	vcvtss2si128Int64(&vals, &ret)

	log.Printf("VCVTSS2SI128INT64 input %v output %d", vals, ret)

	out := number.Int64ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *VCVTSS2SI128INT64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
