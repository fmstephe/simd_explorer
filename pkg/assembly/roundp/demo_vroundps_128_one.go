package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_128_one.s
var assemblyVroundps128One string

//go:embed stub_vroundps_128_one.go
var stubVroundps128One string

type VROUNDPS128ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS128ONE() *VROUNDPS128ONE {
	return &VROUNDPS128ONE{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDPS128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS128ONE) Name() string {
	return "VROUNDPS (128 bit) one"
}

func (v *VROUNDPS128ONE) Description() string {
	return "Round packed singles with imm8=1 (floor)."
}

func (v *VROUNDPS128ONE) Stub() string {
	return stubVroundps128One
}

func (v *VROUNDPS128ONE) Assembly() string {
	return assemblyVroundps128One
}

func (v *VROUNDPS128ONE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundps128One(&vals1, &ret)

	log.Printf("VROUNDPS128ONE vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
