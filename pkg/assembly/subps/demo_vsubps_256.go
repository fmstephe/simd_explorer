package subps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsubps_256.s
var assemblyVsubps256 string

//go:embed stub_vsubps_256.go
var stubVsubps256 string

type VSUBPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSUBPS256() *VSUBPS256 {
	return &VSUBPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSUBPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSUBPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VSUBPS256) Name() string {
	return "VSUBPS (256 bit) "
}

func (v *VSUBPS256) Description() string {
	return "AVX form: subtract packed single-precision floats in YMM, lane-wise (dest - src)."
}

func (v *VSUBPS256) Stub() string {
	return stubVsubps256
}

func (v *VSUBPS256) Assembly() string {
	return assemblyVsubps256
}

func (v *VSUBPS256) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vsubps256(&vals1, &vals2, &ret)

	log.Printf("VSUBPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSUBPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
