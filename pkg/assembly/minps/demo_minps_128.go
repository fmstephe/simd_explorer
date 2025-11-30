package minps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_minps_128.s
var assemblyMinps128 string

//go:embed stub_minps_128.go
var stubMinps128 string

type MINPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewMINPS128() *MINPS128 {
	return &MINPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MINPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *MINPS128) Output() *number.Parameter {
	return v.ret
}

func (v *MINPS128) Name() string {
	return "MINPS XMM (128 bit)"
}

func (v *MINPS128) Description() string {
	return "Compute element-wise minimum of packed single-precision floats in XMM."
}

func (v *MINPS128) Stub() string {
	return stubMinps128
}

func (v *MINPS128) Assembly() string {
	return assemblyMinps128
}

func (v *MINPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	minps128(&vals1, &vals2, &ret)

	log.Printf("MINPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MINPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
