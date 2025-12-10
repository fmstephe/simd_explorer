package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_256_two.s
var assemblyVroundps256Two string

//go:embed stub_vroundps_256_two.go
var stubVroundps256Two string

type VROUNDPS256TWO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS256TWO() *VROUNDPS256TWO {
	return &VROUNDPS256TWO{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VROUNDPS256TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS256TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS256TWO) Name() string {
	return "VROUNDPS (256 bit) two"
}

func (v *VROUNDPS256TWO) Description() string {
	return "Round packed singles with imm8=2 (ceil)."
}

func (v *VROUNDPS256TWO) Stub() string {
	return stubVroundps256Two
}

func (v *VROUNDPS256TWO) Assembly() string {
	return assemblyVroundps256Two
}

func (v *VROUNDPS256TWO) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vroundps256Two(&vals1, &ret)

	log.Printf("VROUNDPS256TWO vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS256TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
