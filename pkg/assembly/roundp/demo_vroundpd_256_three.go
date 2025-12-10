package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundpd_256_three.s
var assemblyVroundpd256Three string

//go:embed stub_vroundpd_256_three.go
var stubVroundpd256Three string

type VROUNDPD256THREE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPD256THREE() *VROUNDPD256THREE {
	return &VROUNDPD256THREE{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VROUNDPD256THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPD256THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPD256THREE) Name() string {
	return "VROUNDPD (256 bit) three"
}

func (v *VROUNDPD256THREE) Description() string {
	return "Round packed doubles with imm8=3 (truncate)."
}

func (v *VROUNDPD256THREE) Stub() string {
	return stubVroundpd256Three
}

func (v *VROUNDPD256THREE) Assembly() string {
	return assemblyVroundpd256Three
}

func (v *VROUNDPD256THREE) Run() {
	vals1 := [4]float64{}
	copy(vals1[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]float64{}

	vroundpd256Three(&vals1, &ret)

	log.Printf("VROUNDPD256THREE vals %v ret %v", vals1, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPD256THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
