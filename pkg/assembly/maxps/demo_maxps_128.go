package maxps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_maxps_128.s
var assemblyMaxps128 string

//go:embed stub_maxps_128.go
var stubMaxps128 string

type MAXPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewMAXPS128() *MAXPS128 {
	return &MAXPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MAXPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *MAXPS128) Output() *number.Parameter {
	return v.ret
}

func (v *MAXPS128) Name() string {
	return "MAXPS XMM (128 bit)"
}

func (v *MAXPS128) Description() string {
	return "Compute element-wise maximum of packed single-precision floats in XMM."
}

func (v *MAXPS128) Stub() string {
	return stubMaxps128
}

func (v *MAXPS128) Assembly() string {
	return assemblyMaxps128
}

func (v *MAXPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	maxps128(&vals1, &vals2, &ret)

	log.Printf("MAXPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MAXPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
