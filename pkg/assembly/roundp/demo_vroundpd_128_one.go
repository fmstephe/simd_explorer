package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_128_one.s
var assemblyVroundpd128One string

//go:embed stub_vroundpd_128_one.go
var stubVroundpd128One string

type VROUNDPD128ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD128ONE() *VROUNDPD128ONE {
	return &VROUNDPD128ONE{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDPD128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD128ONE) Name() string {
	return "VROUNDPD (128 bit) one"
}

func (v *VROUNDPD128ONE) Description() string {
	return "Round packed doubles with imm8=1 (floor)."
}

func (v *VROUNDPD128ONE) Stub() string {
	return stubVroundpd128One
}

func (v *VROUNDPD128ONE) Assembly() string {
	return assemblyVroundpd128One
}

func (v *VROUNDPD128ONE) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundpd128One(&vals1, &ret)

	log.Printf("VROUNDPD128ONE vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
