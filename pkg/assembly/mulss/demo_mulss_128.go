package mulss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_mulss_128.s
var assemblyMulss128 string

//go:embed stub_mulss_128.go
var stubMulss128 string

type MULSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewMULSS128() *MULSS128 {
	return &MULSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MULSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *MULSS128) Output() *number.Parameter {
	return v.ret
}

func (v *MULSS128) Name() string {
	return "MULSS (128 bit) "
}

func (v *MULSS128) Description() string {
	return "Multiply scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *MULSS128) Stub() string {
	return stubMulss128
}

func (v *MULSS128) Assembly() string {
	return assemblyMulss128
}

func (v *MULSS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	mulss128(&vals1, &vals2, &ret)

	log.Printf("MULSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *MULSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
