package subss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_subss_128.s
var assemblySubss128 string

//go:embed stub_subss_128.go
var stubSubss128 string

type SUBSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSUBSS128() *SUBSS128 {
	return &SUBSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SUBSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SUBSS128) Output() *number.Parameter {
	return v.ret
}

func (v *SUBSS128) Name() string {
	return "SUBSS (128 bit) "
}

func (v *SUBSS128) Description() string {
	return "Subtract scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *SUBSS128) Stub() string {
	return stubSubss128
}

func (v *SUBSS128) Assembly() string {
	return assemblySubss128
}

func (v *SUBSS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	subss128(&vals1, &vals2, &ret)

	log.Printf("SUBSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *SUBSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
