package unpcklps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpcklps_256.s
var assemblyVunpcklps256 string

//go:embed stub_vunpcklps_256.go
var stubVunpcklps256 string

type VUNPCKLPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKLPS256() *VUNPCKLPS256 {
	return &VUNPCKLPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VUNPCKLPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKLPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKLPS256) Name() string {
	return "VUNPCKLPS (256 bit) interleave-high"
}

func (v *VUNPCKLPS256) Description() string {
	return "Unpack and interleave high 64-bit halves per 128-bit lane: [a2,b2,a3,b3 | a6,b6,a7,b7]."
}

func (v *VUNPCKLPS256) Stub() string {
	return stubVunpcklps256
}

func (v *VUNPCKLPS256) Assembly() string {
	return assemblyVunpcklps256
}

func (v *VUNPCKLPS256) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vunpcklps256(&vals1, &vals2, &ret)

	log.Printf("VUNPCKLPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VUNPCKLPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
