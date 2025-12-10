package roundp

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vroundps_256_one.s
var assemblyVroundps256One string

//go:embed stub_vroundps_256_one.go
var stubVroundps256One string

type VROUNDPS256ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVROUNDPS256ONE() *VROUNDPS256ONE {
	return &VROUNDPS256ONE{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VROUNDPS256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VROUNDPS256ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VROUNDPS256ONE) Name() string {
	return "VROUNDPS (256 bit) one"
}

func (v *VROUNDPS256ONE) Description() string {
	return "Round packed singles with imm8=1 (floor)."
}

func (v *VROUNDPS256ONE) Stub() string {
	return stubVroundps256One
}

func (v *VROUNDPS256ONE) Assembly() string {
	return assemblyVroundps256One
}

func (v *VROUNDPS256ONE) Run() {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [8]float32{}

	vroundps256One(&vals1, &ret)

	log.Printf("VROUNDPS256ONE vals %v ret %v", vals1, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VROUNDPS256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
