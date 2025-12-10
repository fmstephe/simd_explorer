package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundss_128_zero.s
var assemblyVroundss128Zero string

//go:embed stub_vroundss_128_zero.go
var stubVroundss128Zero string

type VROUNDSS128ZERO struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSS128ZERO() *VROUNDSS128ZERO {
	return &VROUNDSS128ZERO{
		base: number.NewNamedFloatParameter("base", 128, 32),
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDSS128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSS128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSS128ZERO) Name() string {
	return "VROUNDSS (128 bit) zero"
}

func (v *VROUNDSS128ZERO) Description() string {
	return "Round packed singles with imm8=0 (nearest)."
}

func (v *VROUNDSS128ZERO) Stub() string {
	return stubVroundss128Zero
}

func (v *VROUNDSS128ZERO) Assembly() string {
	return assemblyVroundss128Zero
}

func (v *VROUNDSS128ZERO) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundss128Zero(&base, &vals, &ret)

	log.Printf("VROUNDSS128ZERO base %v vals %v ret %v", base, vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSS128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
