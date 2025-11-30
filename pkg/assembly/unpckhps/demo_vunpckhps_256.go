package unpckhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpckhps_256.s
var assemblyVunpckhps256 string

//go:embed stub_vunpckhps_256.go
var stubVunpckhps256 string

type VUNPCKHPS256 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKHPS256() *VUNPCKHPS256 {
	return &VUNPCKHPS256{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VUNPCKHPS256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKHPS256) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKHPS256) Name() string {
	return "VUNPCKHPS (256 bit) interleave-high"
}

func (v *VUNPCKHPS256) Description() string {
	return "Unpack and interleave high 64-bit halves per 128-bit lane: [a2,b2,a3,b3 | a6,b6,a7,b7]."
}

func (v *VUNPCKHPS256) Stub() string {
	return stubVunpckhps256
}

func (v *VUNPCKHPS256) Assembly() string {
	return assemblyVunpckhps256
}

func (v *VUNPCKHPS256) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vunpckhps256(&vals1, &vals2, &ret)

	log.Printf("VUNPCKHPS256 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VUNPCKHPS256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
