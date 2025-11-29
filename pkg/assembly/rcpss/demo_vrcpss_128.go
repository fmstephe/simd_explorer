package rcpss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vrcpss_128.s
var assemblyVrcpss128 string

//go:embed stub_vrcpss_128.go
var stubVrcpss128 string

type VRCPSS128 struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVRCPSS128() *VRCPSS128 {
	return &VRCPSS128{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VRCPSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VRCPSS128) Output() *number.Parameter {
	return v.ret
}

func (v *VRCPSS128) Name() string {
	return "VRCPSS (128 bit) "
}

func (v *VRCPSS128) Description() string {
	return "AVX form: reciprocal estimate of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VRCPSS128) Stub() string {
	return stubVrcpss128
}

func (v *VRCPSS128) Assembly() string {
	return assemblyVrcpss128
}

func (v *VRCPSS128) Run(_ [][]byte) (output []byte) {
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vrcpss128(&vals, &ret)

	log.Printf("VRCPSS128 input %v output %v", vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VRCPSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
