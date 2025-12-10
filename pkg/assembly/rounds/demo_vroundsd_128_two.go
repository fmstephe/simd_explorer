package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundsd_128_two.s
var assemblyVroundsd128Two string

//go:embed stub_vroundsd_128_two.go
var stubVroundsd128Two string

type VROUNDSD128TWO struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSD128TWO() *VROUNDSD128TWO {
	return &VROUNDSD128TWO{
		base: number.NewNamedFloatParameter("base", 128, 64),
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDSD128TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSD128TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSD128TWO) Name() string {
	return "VROUNDSD (128 bit) two"
}

func (v *VROUNDSD128TWO) Description() string {
	return "Round packed doubles with imm8=2 (ceil)."
}

func (v *VROUNDSD128TWO) Stub() string {
	return stubVroundsd128Two
}

func (v *VROUNDSD128TWO) Assembly() string {
	return assemblyVroundsd128Two
}

func (v *VROUNDSD128TWO) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundsd128Two(&base, &vals, &ret)

	log.Printf("VROUNDSD128TWO base %v vals %v ret %v", base, vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSD128TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
