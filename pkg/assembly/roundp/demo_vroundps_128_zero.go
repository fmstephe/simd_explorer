package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_128_zero.s
var assemblyVroundps128Zero string

//go:embed stub_vroundps_128_zero.go
var stubVroundps128Zero string

type VROUNDPS128ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS128ZERO() *VROUNDPS128ZERO {
	return &VROUNDPS128ZERO{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDPS128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS128ZERO) Name() string {
	return "VROUNDPS (128 bit) zero"
}

func (v *VROUNDPS128ZERO) Description() string {
	return "Round packed singles with imm8=0 (nearest)."
}

func (v *VROUNDPS128ZERO) Stub() string {
	return stubVroundps128Zero
}

func (v *VROUNDPS128ZERO) Assembly() string {
	return assemblyVroundps128Zero
}

func (v *VROUNDPS128ZERO) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundps128Zero(&vals1, &ret)

	log.Printf("VROUNDPS128ZERO vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
