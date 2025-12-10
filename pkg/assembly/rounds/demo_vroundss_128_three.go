package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundss_128_three.s
var assemblyVroundss128Three string

//go:embed stub_vroundss_128_three.go
var stubVroundss128Three string

type VROUNDSS128THREE struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSS128THREE() *VROUNDSS128THREE {
	return &VROUNDSS128THREE{
		base: number.NewNamedFloatParameter("base", 128, 32),
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDSS128THREE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSS128THREE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSS128THREE) Name() string {
	return "VROUNDSS (128 bit) three"
}

func (v *VROUNDSS128THREE) Description() string {
	return "Round packed singles with imm8=3 (truncate)."
}

func (v *VROUNDSS128THREE) Stub() string {
	return stubVroundss128Three
}

func (v *VROUNDSS128THREE) Assembly() string {
	return assemblyVroundss128Three
}

func (v *VROUNDSS128THREE) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundss128Three(&base, &vals, &ret)

	log.Printf("VROUNDSS128THREE base %v vals %v ret %v", base, vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSS128THREE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
