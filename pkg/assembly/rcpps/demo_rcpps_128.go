package rcpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_rcpps_128.s
var assemblyRcpps128 string

//go:embed stub_rcpps_128.go
var stubRcpps128 string

type RCPPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewRCPPS128() *RCPPS128 {
	return &RCPPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *RCPPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *RCPPS128) Output() *number.Parameter {
	return v.ret
}

func (v *RCPPS128) Name() string {
	return "RCPPS XMM (128 bit)"
}

func (v *RCPPS128) Description() string {
	return "Compute reciprocal estimate of packed single-precision floats in XMM, lane-wise."
}

func (v *RCPPS128) Stub() string {
	return stubRcpps128
}

func (v *RCPPS128) Assembly() string {
	return assemblyRcpps128
}

func (v *RCPPS128) Run() {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	rcpps128(&vals, &ret)

	log.Printf("RCPPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *RCPPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
