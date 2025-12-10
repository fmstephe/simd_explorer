package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_128_two.s
var assemblyVroundps128Two string

//go:embed stub_vroundps_128_two.go
var stubVroundps128Two string

type VROUNDPS128TWO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS128TWO() *VROUNDPS128TWO {
	return &VROUNDPS128TWO{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDPS128TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS128TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS128TWO) Name() string {
	return "VROUNDPS (128 bit) two"
}

func (v *VROUNDPS128TWO) Description() string {
	return "Round packed singles with imm8=2 (ceil)."
}

func (v *VROUNDPS128TWO) Stub() string {
	return stubVroundps128Two
}

func (v *VROUNDPS128TWO) Assembly() string {
	return assemblyVroundps128Two
}

func (v *VROUNDPS128TWO) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundps128Two(&vals1, &ret)

	log.Printf("VROUNDPS128TWO vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS128TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
