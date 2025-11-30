package minss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_minss_128.s
var assemblyMinss128 string

//go:embed stub_minss_128.go
var stubMinss128 string

type MINSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewMINSS128() *MINSS128 {
	return &MINSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MINSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *MINSS128) Output() *number.Parameter {
	return v.ret
}

func (v *MINSS128) Name() string {
	return "MINSS (128 bit) "
}

func (v *MINSS128) Description() string {
	return "Compute minimum of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *MINSS128) Stub() string {
	return stubMinss128
}

func (v *MINSS128) Assembly() string {
	return assemblyMinss128
}

func (v *MINSS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	minss128(&vals1, &vals2, &ret)

	log.Printf("MINSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MINSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
