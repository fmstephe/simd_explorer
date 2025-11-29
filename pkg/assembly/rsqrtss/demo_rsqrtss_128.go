package rsqrtss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rsqrtss_128.s
var assemblyRsqrtss128 string

//go:embed stub_rsqrtss_128.go
var stubRsqrtss128 string

type RSQRTSS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewRSQRTSS128() *RSQRTSS128 {
	return &RSQRTSS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *RSQRTSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *RSQRTSS128) Output() *number.Parameter {
	return v.ret
}

func (v *RSQRTSS128) Name() string {
	return "RSQRTSS (128 bit) "
}

func (v *RSQRTSS128) Description() string {
	return "Compute reciprocal square root estimate of scalar single-precision (lane 0); upper lanes pass through."
}

func (v *RSQRTSS128) Stub() string {
	return stubRsqrtss128
}

func (v *RSQRTSS128) Assembly() string {
	return assemblyRsqrtss128
}

func (v *RSQRTSS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	rsqrtss128(&vals, &ret)

	log.Printf("RSQRTSS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *RSQRTSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
