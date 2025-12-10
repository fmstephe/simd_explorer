package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_256_two.s
var assemblyVroundpd256Two string

//go:embed stub_vroundpd_256_two.go
var stubVroundpd256Two string

type VROUNDPD256TWO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD256TWO() *VROUNDPD256TWO {
	return &VROUNDPD256TWO{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VROUNDPD256TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD256TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD256TWO) Name() string {
	return "VROUNDPD (256 bit) two"
}

func (v *VROUNDPD256TWO) Description() string {
	return "Round packed doubles with imm8=2 (ceil)."
}

func (v *VROUNDPD256TWO) Stub() string {
	return stubVroundpd256Two
}

func (v *VROUNDPD256TWO) Assembly() string {
	return assemblyVroundpd256Two
}

func (v *VROUNDPD256TWO) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]float64{}

	vroundpd256Two(&vals1, &ret)

	log.Printf("VROUNDPD256TWO vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD256TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
