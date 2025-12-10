package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_128_two.s
var assemblyVroundpd128Two string

//go:embed stub_vroundpd_128_two.go
var stubVroundpd128Two string

type VROUNDPD128TWO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD128TWO() *VROUNDPD128TWO {
	return &VROUNDPD128TWO{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDPD128TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD128TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD128TWO) Name() string {
	return "VROUNDPD (128 bit) two"
}

func (v *VROUNDPD128TWO) Description() string {
	return "Round packed doubles with imm8=2 (ceil)."
}

func (v *VROUNDPD128TWO) Stub() string {
	return stubVroundpd128Two
}

func (v *VROUNDPD128TWO) Assembly() string {
	return assemblyVroundpd128Two
}

func (v *VROUNDPD128TWO) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundpd128Two(&vals1, &ret)

	log.Printf("VROUNDPD128TWO vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD128TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
