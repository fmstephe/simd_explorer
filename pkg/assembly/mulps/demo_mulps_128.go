package mulps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_mulps_128.s
var assemblyMulps128 string

//go:embed stub_mulps_128.go
var stubMulps128 string

type MULPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewMULPS128() *MULPS128 {
	return &MULPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *MULPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *MULPS128) Output() *number.Parameter {
	return v.ret
}

func (v *MULPS128) Name() string {
	return "MULPS (128 bit) "
}

func (v *MULPS128) Description() string {
	return "Multiply packed single-precision floats in XMM, lane-wise."
}

func (v *MULPS128) Stub() string {
	return stubMulps128
}

func (v *MULPS128) Assembly() string {
	return assemblyMulps128
}

func (v *MULPS128) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	mulps128(&vals1, &vals2, &ret)

	log.Printf("MULPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *MULPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
