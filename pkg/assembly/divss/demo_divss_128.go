package divss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_divss_128.s
var assemblyDivss128 string

//go:embed stub_divss_128.go
var stubDivss128 string

type DIVSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewDIVSS128() *DIVSS128 {
	return &DIVSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *DIVSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *DIVSS128) Output() *number.Parameter {
	return v.ret
}

func (v *DIVSS128) Name() string {
	return "DIVSS (128 bit) "
}

func (v *DIVSS128) Description() string {
	return "Divide scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *DIVSS128) Stub() string {
	return stubDivss128
}

func (v *DIVSS128) Assembly() string {
	return assemblyDivss128
}

func (v *DIVSS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	divss128(&vals1, &vals2, &ret)

	log.Printf("DIVSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *DIVSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
