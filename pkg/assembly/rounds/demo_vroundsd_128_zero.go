package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundsd_128_zero.s
var assemblyVroundsd128Zero string

//go:embed stub_vroundsd_128_zero.go
var stubVroundsd128Zero string

type VROUNDSD128ZERO struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSD128ZERO() *VROUNDSD128ZERO {
	return &VROUNDSD128ZERO{
		base: number.NewNamedFloatParameter("base", 128, 64),
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDSD128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSD128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSD128ZERO) Name() string {
	return "VROUNDSD (128 bit) zero"
}

func (v *VROUNDSD128ZERO) Description() string {
	return "Round packed doubles with imm8=0 (nearest)."
}

func (v *VROUNDSD128ZERO) Stub() string {
	return stubVroundsd128Zero
}

func (v *VROUNDSD128ZERO) Assembly() string {
	return assemblyVroundsd128Zero
}

func (v *VROUNDSD128ZERO) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundsd128Zero(&base, &vals, &ret)

	log.Printf("VROUNDSD128ZERO base %v vals %v ret %v", base, vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSD128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
