package divss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vdivss_128.s
var assemblyVdivss128 string

//go:embed stub_vdivss_128.go
var stubVdivss128 string

type VDIVSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVDIVSS128() *VDIVSS128 {
	return &VDIVSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VDIVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VDIVSS128) Output() *number.Parameter {
	return v.ret
}

func (v *VDIVSS128) Name() string {
	return "VDIVSS (128 bit) "
}

func (v *VDIVSS128) Description() string {
	return "AVX form: divide scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *VDIVSS128) Stub() string {
	return stubVdivss128
}

func (v *VDIVSS128) Assembly() string {
	return assemblyVdivss128
}

func (v *VDIVSS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vdivss128(&vals1, &vals2, &ret)

	log.Printf("VDIVSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VDIVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
