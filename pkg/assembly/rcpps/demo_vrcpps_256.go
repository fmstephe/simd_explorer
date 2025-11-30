package rcpps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrcpps_256.s
var assemblyVrcpps256 string

//go:embed stub_vrcpps_256.go
var stubVrcpps256 string

type VRCPPS256 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVRCPPS256() *VRCPPS256 {
	return &VRCPPS256{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VRCPPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VRCPPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VRCPPS256) Name() string {
	return "VRCPPS YMM (256 bit)"
}

func (v *VRCPPS256) Description() string {
	return "AVX form: reciprocal estimate of packed single-precision floats in YMM, lane-wise."
}

func (v *VRCPPS256) Stub() string {
	return stubVrcpps256
}

func (v *VRCPPS256) Assembly() string {
	return assemblyVrcpps256
}

func (v *VRCPPS256) Run() (output []byte) {
	vals := [8]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vrcpps256(&vals, &ret)

	log.Printf("VRCPPS256 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VRCPPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
