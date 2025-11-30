package maxss

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_maxss_128.s
var assemblyMaxss128 string

//go:embed stub_maxss_128.go
var stubMaxss128 string

type MAXSS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewMAXSS128() *MAXSS128 {
	return &MAXSS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MAXSS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *MAXSS128) Output() *number.Parameter {
	return v.ret
}

func (v *MAXSS128) Name() string {
	return "MAXSS (128 bit) "
}

func (v *MAXSS128) Description() string {
	return "Compute maximum of scalar single-precision (lane 0); upper lanes pass through from the first operand."
}

func (v *MAXSS128) Stub() string {
	return stubMaxss128
}

func (v *MAXSS128) Assembly() string {
	return assemblyMaxss128
}

func (v *MAXSS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	maxss128(&vals1, &vals2, &ret)

	log.Printf("MAXSS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MAXSS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
