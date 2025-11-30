package rcpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rcpss_128.s
var assemblyRcpss128 string

//go:embed stub_rcpss_128.go
var stubRcpss128 string

type RCPSS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewRCPSS128() *RCPSS128 {
	return &RCPSS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *RCPSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *RCPSS128) Output() *number.Parameter {
	return v.ret
}

func (v *RCPSS128) Name() string {
	return "RCPSS (128 bit) "
}

func (v *RCPSS128) Description() string {
	return "Compute reciprocal estimate of scalar single-precision (lane 0); upper lanes pass through."
}

func (v *RCPSS128) Stub() string {
	return stubRcpss128
}

func (v *RCPSS128) Assembly() string {
	return assemblyRcpss128
}

func (v *RCPSS128) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	rcpss128(&vals, &ret)

	log.Printf("RCPSS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *RCPSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
