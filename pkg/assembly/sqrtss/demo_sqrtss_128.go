package sqrtss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_sqrtss_128.s
var assemblySqrtss128 string

//go:embed stub_sqrtss_128.go
var stubSqrtss128 string

type SQRTSS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewSQRTSS128() *SQRTSS128 {
	return &SQRTSS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SQRTSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *SQRTSS128) Output() *number.Parameter {
	return v.ret
}

func (v *SQRTSS128) Name() string {
	return "SQRTSS (128 bit) "
}

func (v *SQRTSS128) Description() string {
	return "Compute square root of scalar single-precision (lane 0); upper lanes pass through."
}

func (v *SQRTSS128) Stub() string {
	return stubSqrtss128
}

func (v *SQRTSS128) Assembly() string {
	return assemblySqrtss128
}

func (v *SQRTSS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	sqrtss128(&vals, &ret)

	log.Printf("SQRTSS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SQRTSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
