package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundss_128_two.s
var assemblyVroundss128Two string

//go:embed stub_vroundss_128_two.go
var stubVroundss128Two string

type VROUNDSS128TWO struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSS128TWO() *VROUNDSS128TWO {
	return &VROUNDSS128TWO{
		base: number.NewNamedFloatParameter("base", 128, 32),
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDSS128TWO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSS128TWO) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSS128TWO) Name() string {
	return "VROUNDSS (128 bit) two"
}

func (v *VROUNDSS128TWO) Description() string {
	return "Round packed singles with imm8=2 (ceil)."
}

func (v *VROUNDSS128TWO) Stub() string {
	return stubVroundss128Two
}

func (v *VROUNDSS128TWO) Assembly() string {
	return assemblyVroundss128Two
}

func (v *VROUNDSS128TWO) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundss128Two(&base, &vals, &ret)

	log.Printf("VROUNDSS128TWO base %v vals %v ret %v", base, vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSS128TWO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
