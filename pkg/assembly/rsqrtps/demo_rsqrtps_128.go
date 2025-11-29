package rsqrtps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rsqrtps_128.s
var assemblyRsqrtps128 string

//go:embed stub_rsqrtps_128.go
var stubRsqrtps128 string

type RSQRTPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewRSQRTPS128() *RSQRTPS128 {
	return &RSQRTPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *RSQRTPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *RSQRTPS128) Output() *number.Parameter {
	return v.ret
}

func (v *RSQRTPS128) Name() string {
	return "RSQRTPS XMM (128 bit)"
}

func (v *RSQRTPS128) Description() string {
	return "Compute reciprocal square root estimate of packed single-precision floats in XMM, lane-wise."
}

func (v *RSQRTPS128) Stub() string {
	return stubRsqrtps128
}

func (v *RSQRTPS128) Assembly() string {
	return assemblyRsqrtps128
}

func (v *RSQRTPS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	rsqrtps128(&vals, &ret)

	log.Printf("RSQRTPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *RSQRTPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
