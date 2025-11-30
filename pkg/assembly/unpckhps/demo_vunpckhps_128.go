package unpckhps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vunpckhps_128.s
var assemblyVunpckhps128 string

//go:embed stub_vunpckhps_128.go
var stubVunpckhps128 string

type VUNPCKHPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVUNPCKHPS128() *VUNPCKHPS128 {
	return &VUNPCKHPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VUNPCKHPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VUNPCKHPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VUNPCKHPS128) Name() string {
	return "VUNPCKHPS (128 bit) interleave-high"
}

func (v *VUNPCKHPS128) Description() string {
	return "Unpack and interleave high 64-bit halves: dst = [a2,b2,a3,b3]."
}

func (v *VUNPCKHPS128) Stub() string {
	return stubVunpckhps128
}

func (v *VUNPCKHPS128) Assembly() string {
	return assemblyVunpckhps128
}

func (v *VUNPCKHPS128) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vunpckhps128(&vals1, &vals2, &ret)

	log.Printf("VUNPCKHPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VUNPCKHPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
