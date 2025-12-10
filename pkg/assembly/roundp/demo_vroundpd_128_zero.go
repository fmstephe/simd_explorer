package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_128_zero.s
var assemblyVroundpd128Zero string

//go:embed stub_vroundpd_128_zero.go
var stubVroundpd128Zero string

type VROUNDPD128ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD128ZERO() *VROUNDPD128ZERO {
	return &VROUNDPD128ZERO{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDPD128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD128ZERO) Name() string {
	return "VROUNDPD (128 bit) zero"
}

func (v *VROUNDPD128ZERO) Description() string {
	return "Round packed doubles with imm8=0 (nearest)."
}

func (v *VROUNDPD128ZERO) Stub() string {
	return stubVroundpd128Zero
}

func (v *VROUNDPD128ZERO) Assembly() string {
	return assemblyVroundpd128Zero
}

func (v *VROUNDPD128ZERO) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundpd128Zero(&vals1, &ret)

	log.Printf("VROUNDPD128ZERO vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
