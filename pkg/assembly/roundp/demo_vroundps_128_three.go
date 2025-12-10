package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_128_three.s
var assemblyVroundps128Three string

//go:embed stub_vroundps_128_three.go
var stubVroundps128Three string

type VROUNDPS128THREE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS128THREE() *VROUNDPS128THREE {
	return &VROUNDPS128THREE{
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDPS128THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS128THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS128THREE) Name() string {
	return "VROUNDPS (128 bit) three"
}

func (v *VROUNDPS128THREE) Description() string {
	return "Round packed singles with imm8=3 (truncate)."
}

func (v *VROUNDPS128THREE) Stub() string {
	return stubVroundps128Three
}

func (v *VROUNDPS128THREE) Assembly() string {
	return assemblyVroundps128Three
}

func (v *VROUNDPS128THREE) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundps128Three(&vals1, &ret)

	log.Printf("VROUNDPS128THREE vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS128THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
