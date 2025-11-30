package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_threes.s
var assemblyShufps128Threes string

//go:embed stub_shufps_128_threes.go
var stubShufps128Threes string

type SHUFPS128THREES struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSHUFPS128THREES() *SHUFPS128THREES {
	return &SHUFPS128THREES{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SHUFPS128THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SHUFPS128THREES) Output() *number.Parameter {
	return v.ret
}

func (v *SHUFPS128THREES) Name() string {
	return "SHUFPS (128 bit) threes"
}

func (v *SHUFPS128THREES) Description() string {
	return "SHUFPS imm8=0xFF: dst = [a3,a3,b3,b3]"
}

func (v *SHUFPS128THREES) Stub() string {
	return stubShufps128Threes
}

func (v *SHUFPS128THREES) Assembly() string {
	return assemblyShufps128Threes
}

func (v *SHUFPS128THREES) Run() (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	shufps128Threes(&vals1, &vals2, &ret)

	log.Printf("SHUFPS128THREES input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SHUFPS128THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
