package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_256_three.s
var assemblyVroundps256Three string

//go:embed stub_vroundps_256_three.go
var stubVroundps256Three string

type VROUNDPS256THREE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS256THREE() *VROUNDPS256THREE {
	return &VROUNDPS256THREE{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VROUNDPS256THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS256THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS256THREE) Name() string {
	return "VROUNDPS (256 bit) three"
}

func (v *VROUNDPS256THREE) Description() string {
	return "Round packed singles with imm8=3 (truncate)."
}

func (v *VROUNDPS256THREE) Stub() string {
	return stubVroundps256Three
}

func (v *VROUNDPS256THREE) Assembly() string {
	return assemblyVroundps256Three
}

func (v *VROUNDPS256THREE) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vroundps256Three(&vals1, &ret)

	log.Printf("VROUNDPS256THREE vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS256THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
