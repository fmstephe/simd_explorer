package orps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vorps_128.s
var assemblyVorps128 string

//go:embed stub_vorps_128.go
var stubVorps128 string

type VORPS128 struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVORPS128() *VORPS128 {
	return &VORPS128{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VORPS128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VORPS128) Output() *number.Parameter {
	return v.ret
}

func (v *VORPS128) Name() string {
	return "VORPS (128 bit)"
}

func (v *VORPS128) Description() string {
	return "Bitwise AND of packed single-precision values; output shown as 32-bit hex lanes."
}

func (v *VORPS128) Stub() string {
	return stubVorps128
}

func (v *VORPS128) Assembly() string {
	return assemblyVorps128
}

func (v *VORPS128) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vorps128(&vals1, &vals2, &ret)

	log.Printf("VORPS128 input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VORPS128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
