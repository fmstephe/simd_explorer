package unpcklps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpcklps_128.s
var assemblyVunpcklps128 string

//go:embed stub_vunpcklps_128.go
var stubVunpcklps128 string

type VUNPCKLPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKLPS128() *VUNPCKLPS128 {
	return &VUNPCKLPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VUNPCKLPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKLPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKLPS128) Name() string {
	return "VUNPCKLPS (128 bit) interleave-high"
}

func (v *VUNPCKLPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *VUNPCKLPS128) Stub() string {
	return stubVunpcklps128
}

func (v *VUNPCKLPS128) Assembly() string {
	return assemblyVunpcklps128
}

func (v *VUNPCKLPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vunpcklps128(&vals1, &vals2, &ret)

	log.Printf("VUNPCKLPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VUNPCKLPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
