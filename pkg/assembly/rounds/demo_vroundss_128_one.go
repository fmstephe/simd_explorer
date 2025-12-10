package rounds

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundss_128_one.s
var assemblyVroundss128One string

//go:embed stub_vroundss_128_one.go
var stubVroundss128One string

type VROUNDSS128ONE struct {
	base *number.Parameter
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDSS128ONE() *VROUNDSS128ONE {
	return &VROUNDSS128ONE{
		base: number.NewNamedFloatParameter("base", 128, 32),
		vals: number.NewNamedFloatParameter("vals", 128, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VROUNDSS128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.base,
		v.vals,
	}
}

func (v *VROUNDSS128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDSS128ONE) Name() string {
	return "VROUNDSS (128 bit) one"
}

func (v *VROUNDSS128ONE) Description() string {
	return "Round packed singles with imm8=1 (floor)."
}

func (v *VROUNDSS128ONE) Stub() string {
	return stubVroundss128One
}

func (v *VROUNDSS128ONE) Assembly() string {
	return assemblyVroundss128One
}

func (v *VROUNDSS128ONE) Run() {
	base := [4]float32{}
	copy(base[:], number.ToFloat32Slice(v.base.FlatData()))
	vals := [4]float32{}
	copy(vals[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vroundss128One(&base, &vals, &ret)

	log.Printf("VROUNDSS128ONE base %v vals %v ret %v", base, vals, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDSS128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
