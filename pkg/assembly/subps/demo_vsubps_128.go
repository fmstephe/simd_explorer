package subps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vsubps_128.s
var assemblyVsubps128 string

//go:embed stub_vsubps_128.go
var stubVsubps128 string

type VSUBPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSUBPS128() *VSUBPS128 {
	return &VSUBPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSUBPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSUBPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VSUBPS128) Name() string {
	return "VSUBPS (128 bit) "
}

func (v *VSUBPS128) Description() string {
	return "AVX form: subtract packed single-precision floats in XMM, lane-wise (dest - src)."
}

func (v *VSUBPS128) Stub() string {
	return stubVsubps128
}

func (v *VSUBPS128) Assembly() string {
	return assemblyVsubps128
}

func (v *VSUBPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vsubps128(&vals1, &vals2, &ret)

	log.Printf("VSUBPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSUBPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
