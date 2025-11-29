package addss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vaddss_128.s
var assemblyVaddss128 string

//go:embed stub_vaddss_128.go
var stubVaddss128 string

type VADDSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVADDSS128() *VADDSS128 {
	return &VADDSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VADDSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VADDSS128) Output() *number.Parameter {
	return v.ret
}

func (v *VADDSS128) Name() string {
	return "VADDSS (128 bit) "
}

func (v *VADDSS128) Description() string {
	return "AVX form: add scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VADDSS128) Stub() string {
	return stubVaddss128
}

func (v *VADDSS128) Assembly() string {
	return assemblyVaddss128
}

func (v *VADDSS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	addss128(&vals1, &vals2, &ret)

	log.Printf("VADDSS128 input %v %v output %v", vals1, vals2, ret)

	retSlc := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(retSlc)
	return retSlc
}

func (v *VADDSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
