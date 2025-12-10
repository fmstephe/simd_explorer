package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundsd_128_one.s
var assemblyVroundsd128One string

//go:embed stub_vroundsd_128_one.go
var stubVroundsd128One string

type VROUNDSD128ONE struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSD128ONE() *VROUNDSD128ONE {
	return &VROUNDSD128ONE{
		base: number.NewNamedFloatParameter("base", 128, 64),
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDSD128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSD128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSD128ONE) Name() string {
	return "VROUNDSD (128 bit) one"
}

func (v *VROUNDSD128ONE) Description() string {
	return "Round packed doubles with imm8=1 (floor)."
}

func (v *VROUNDSD128ONE) Stub() string {
	return stubVroundsd128One
}

func (v *VROUNDSD128ONE) Assembly() string {
	return assemblyVroundsd128One
}

func (v *VROUNDSD128ONE) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundsd128One(&base, &vals, &ret)

	log.Printf("VROUNDSD128ONE base %v vals %v ret %v", base, vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSD128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
