package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundsd_128_three.s
var assemblyVroundsd128Three string

//go:embed stub_vroundsd_128_three.go
var stubVroundsd128Three string

type VROUNDSD128THREE struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSD128THREE() *VROUNDSD128THREE {
	return &VROUNDSD128THREE{
		base: number.NewNamedFloatParameter("base", 128, 64),
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VROUNDSD128THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSD128THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSD128THREE) Name() string {
	return "VROUNDSD (128 bit) three"
}

func (v *VROUNDSD128THREE) Description() string {
	return "Round packed doubles with imm8=3 (truncate)."
}

func (v *VROUNDSD128THREE) Stub() string {
	return stubVroundsd128Three
}

func (v *VROUNDSD128THREE) Assembly() string {
	return assemblyVroundsd128Three
}

func (v *VROUNDSD128THREE) Run() {
	base := [2]float64{}
	copy(base[:], number.ToFloat64Slice(v.base.FlatData()))
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vroundsd128Three(&base, &vals, &ret)

	log.Printf("VROUNDSD128THREE base %v vals %v ret %v", base, vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSD128THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
