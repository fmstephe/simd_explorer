package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_256_one.s
var assemblyVroundpd256One string

//go:embed stub_vroundpd_256_one.go
var stubVroundpd256One string

type VROUNDPD256ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD256ONE() *VROUNDPD256ONE {
	return &VROUNDPD256ONE{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VROUNDPD256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD256ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD256ONE) Name() string {
	return "VROUNDPD (256 bit) one"
}

func (v *VROUNDPD256ONE) Description() string {
	return "Round packed doubles with imm8=1 (floor)."
}

func (v *VROUNDPD256ONE) Stub() string {
	return stubVroundpd256One
}

func (v *VROUNDPD256ONE) Assembly() string {
	return assemblyVroundpd256One
}

func (v *VROUNDPD256ONE) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]float64{}

	vroundpd256One(&vals1, &ret)

	log.Printf("VROUNDPD256ONE vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
