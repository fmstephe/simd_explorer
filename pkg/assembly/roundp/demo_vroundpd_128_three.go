package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_128_three.s
var assemblyVroundpd128Three string

//go:embed stub_vroundpd_128_three.go
var stubVroundpd128Three string

type VROUNDPD128THREE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD128THREE() *VROUNDPD128THREE {
	return &VROUNDPD128THREE{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDPD128THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD128THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD128THREE) Name() string {
	return "VROUNDPD (128 bit) three"
}

func (v *VROUNDPD128THREE) Description() string {
	return "Round packed doubles with imm8=3 (truncate)."
}

func (v *VROUNDPD128THREE) Stub() string {
	return stubVroundpd128Three
}

func (v *VROUNDPD128THREE) Assembly() string {
	return assemblyVroundpd128Three
}

func (v *VROUNDPD128THREE) Run() {
	vals1 := [2]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundpd128Three(&vals1, &ret)

	log.Printf("VROUNDPD128THREE vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD128THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
