package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_256_zero.s
var assemblyVroundpd256Zero string

//go:embed stub_vroundpd_256_zero.go
var stubVroundpd256Zero string

type VROUNDPD256ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD256ZERO() *VROUNDPD256ZERO {
	return &VROUNDPD256ZERO{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VROUNDPD256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD256ZERO) Name() string {
	return "VROUNDPD (256 bit) zero"
}

func (v *VROUNDPD256ZERO) Description() string {
	return "Round packed doubles with imm8=0 (nearest)."
}

func (v *VROUNDPD256ZERO) Stub() string {
	return stubVroundpd256Zero
}

func (v *VROUNDPD256ZERO) Assembly() string {
	return assemblyVroundpd256Zero
}

func (v *VROUNDPD256ZERO) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]float64{}

	vroundpd256Zero(&vals1, &ret)

	log.Printf("VROUNDPD256ZERO vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
