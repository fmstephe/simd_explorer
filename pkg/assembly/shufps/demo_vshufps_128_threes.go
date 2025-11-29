package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_threes.s
var assemblyVshufps128Threes string

//go:embed stub_vshufps_128_threes.go
var stubVshufps128Threes string

type VSHUFPS128THREES struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS128THREES() *VSHUFPS128THREES {
	return &VSHUFPS128THREES{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSHUFPS128THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS128THREES) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS128THREES) Name() string {
	return "VSHUFPS (128 bit) threes"
}

func (v *VSHUFPS128THREES) Description() string {
	return "VSHUFPS imm8=0xFF: dst = [a3,a3,b3,b3]"
}

func (v *VSHUFPS128THREES) Stub() string {
	return stubVshufps128Threes
}

func (v *VSHUFPS128THREES) Assembly() string {
	return assemblyVshufps128Threes
}

func (v *VSHUFPS128THREES) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vshufps128Threes(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS128THREES input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS128THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
