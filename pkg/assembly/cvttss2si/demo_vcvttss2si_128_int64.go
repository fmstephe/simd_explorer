package cvttss2si

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvttss2si_128_int64.s
var assemblyVcvttss2si128Int64 string

//go:embed stub_vcvttss2si_128_int64.go
var stubVcvttss2si128Int64 string

type VCVTTSS2SI128INT64 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTTSS2SI128INT64() *VCVTTSS2SI128INT64 {
	return &VCVTTSS2SI128INT64{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedIntParameter("ret", 64, 64, 10),
	}
}

func (v *VCVTTSS2SI128INT64) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTTSS2SI128INT64) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTTSS2SI128INT64) Name() string {
	return "VCVTTSS2SI (128 bit) int64"
}

func (v *VCVTTSS2SI128INT64) Description() string {
	return "Truncate scalar single-precision (lowest lane) to signed 64-bit integer."
}

func (v *VCVTTSS2SI128INT64) Stub() string {
	return stubVcvttss2si128Int64
}

func (v *VCVTTSS2SI128INT64) Assembly() string {
	return assemblyVcvttss2si128Int64
}

func (v *VCVTTSS2SI128INT64) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	var ret int64

	vcvttss2si128Int64(&vals, &ret)

	log.Printf("VCVTTSS2SI128INT64 input %v output %d", vals, ret)

	out := number.Int64ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *VCVTTSS2SI128INT64) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
