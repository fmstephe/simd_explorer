package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_threes.s
var assemblyVshufps256Threes string

//go:embed stub_vshufps_256_threes.go
var stubVshufps256Threes string

type VSHUFPS256THREES struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS256THREES() *VSHUFPS256THREES {
	return &VSHUFPS256THREES{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSHUFPS256THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS256THREES) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS256THREES) Name() string {
	return "VSHUFPS (256 bit) threes"
}

func (v *VSHUFPS256THREES) Description() string {
	return "VSHUFPS imm8=0xFF per 128-bit lane: dst = [a3,a3,b3,b3 | a7,a7,b7,b7]"
}

func (v *VSHUFPS256THREES) Stub() string {
	return stubVshufps256Threes
}

func (v *VSHUFPS256THREES) Assembly() string {
	return assemblyVshufps256Threes
}

func (v *VSHUFPS256THREES) Run() (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vshufps256Threes(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS256THREES input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS256THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
