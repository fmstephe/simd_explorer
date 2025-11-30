package unpcklps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_unpcklps_128.s
var assemblyUnpcklps128 string

//go:embed stub_unpcklps_128.go
var stubUnpcklps128 string

type UNPCKLPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewUNPCKLPS128() *UNPCKLPS128 {
	return &UNPCKLPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *UNPCKLPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *UNPCKLPS128) Output() *number.Parameter {
	return v.ret
}

func (v *UNPCKLPS128) Name() string {
	return "UNPCKLPS (128 bit) interleave-high"
}

func (v *UNPCKLPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *UNPCKLPS128) Stub() string {
	return stubUnpcklps128
}

func (v *UNPCKLPS128) Assembly() string {
	return assemblyUnpcklps128
}

func (v *UNPCKLPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	unpcklps128(&vals1, &vals2, &ret)

	log.Printf("UNPCKLPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *UNPCKLPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
