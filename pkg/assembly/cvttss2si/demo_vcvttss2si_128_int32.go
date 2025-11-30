package cvttss2si

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vcvttss2si_128_int32.s
var assemblyVcvttss2si128Int32 string

//go:embed stub_vcvttss2si_128_int32.go
var stubVcvttss2si128Int32 string

type VCVTTSS2SI128INT32 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVCVTTSS2SI128INT32() *VCVTTSS2SI128INT32 {
	return &VCVTTSS2SI128INT32{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedIntParameter("ret", 32, 32, 10),
	}
}

func (v *VCVTTSS2SI128INT32) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VCVTTSS2SI128INT32) Output() *number.Parameter {
	return v.ret
}

func (v *VCVTTSS2SI128INT32) Name() string {
	return "VCVTTSS2SI (128 bit) int32"
}

func (v *VCVTTSS2SI128INT32) Description() string {
	return "Truncate scalar single-precision (lowest lane) to signed 32-bit integer."
}

func (v *VCVTTSS2SI128INT32) Stub() string {
	return stubVcvttss2si128Int32
}

func (v *VCVTTSS2SI128INT32) Assembly() string {
	return assemblyVcvttss2si128Int32
}

func (v *VCVTTSS2SI128INT32) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	var ret int32

	vcvttss2si128Int32(&vals, &ret)

	log.Printf("VCVTTSS2SI128INT32 input %v output %d", vals, ret)

	out := number.Int32ToBytes(ret)
	v.ret.SetData(out)
	return out
}

func (v *VCVTTSS2SI128INT32) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
