package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_256_zero.s
var assemblyVroundps256Zero string

//go:embed stub_vroundps_256_zero.go
var stubVroundps256Zero string

type VROUNDPS256ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS256ZERO() *VROUNDPS256ZERO {
	return &VROUNDPS256ZERO{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VROUNDPS256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS256ZERO) Name() string {
	return "VROUNDPS (256 bit) zero"
}

func (v *VROUNDPS256ZERO) Description() string {
	return "Round packed singles with imm8=0 (nearest)."
}

func (v *VROUNDPS256ZERO) Stub() string {
	return stubVroundps256Zero
}

func (v *VROUNDPS256ZERO) Assembly() string {
	return assemblyVroundps256Zero
}

func (v *VROUNDPS256ZERO) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vroundps256Zero(&vals1, &ret)

	log.Printf("VROUNDPS256ZERO vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
