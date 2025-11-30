package rcpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrcpps_128.s
var assemblyVrcpps128 string

//go:embed stub_vrcpps_128.go
var stubVrcpps128 string

type VRCPPS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVRCPPS128() *VRCPPS128 {
	return &VRCPPS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VRCPPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VRCPPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VRCPPS128) Name() string {
	return "VRCPPS XMM (128 bit)"
}

func (v *VRCPPS128) Description() string {
	return "AVX form: reciprocal estimate of packed single-precision floats in XMM, lane-wise."
}

func (v *VRCPPS128) Stub() string {
	return stubVrcpps128
}

func (v *VRCPPS128) Assembly() string {
	return assemblyVrcpps128
}

func (v *VRCPPS128) Run() (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vrcpps128(&vals, &ret)

	log.Printf("VRCPPS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VRCPPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
