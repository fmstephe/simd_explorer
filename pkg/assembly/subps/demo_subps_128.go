package subps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_subps_128.s
var assemblySubps128 string

//go:embed stub_subps_128.go
var stubSubps128 string

type SUBPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSUBPS128() *SUBPS128 {
	return &SUBPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SUBPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SUBPS128) Output() *number.Parameter {
	return v.ret
}

func (v *SUBPS128) Name() string {
	return "SUBPS (128 bit) "
}

func (v *SUBPS128) Description() string {
	return "Subtract packed single-precision floats in XMM, lane-wise (dest - src)."
}

func (v *SUBPS128) Stub() string {
	return stubSubps128
}

func (v *SUBPS128) Assembly() string {
	return assemblySubps128
}

func (v *SUBPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	subps128(&vals1, &vals2, &ret)

	log.Printf("SUBPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SUBPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
