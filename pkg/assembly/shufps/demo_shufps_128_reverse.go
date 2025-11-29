package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_reverse.s
var assemblyShufps128Reverse string

//go:embed stub_shufps_128_reverse.go
var stubShufps128Reverse string

type SHUFPS128REVERSE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSHUFPS128REVERSE() *SHUFPS128REVERSE {
	return &SHUFPS128REVERSE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SHUFPS128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SHUFPS128REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *SHUFPS128REVERSE) Name() string {
	return "SHUFPS (128 bit) reverse"
}

func (v *SHUFPS128REVERSE) Description() string {
	return "SHUFPS imm8=0x1B: dst = [a3,a2,b1,b0]"
}

func (v *SHUFPS128REVERSE) Stub() string {
	return stubShufps128Reverse
}

func (v *SHUFPS128REVERSE) Assembly() string {
	return assemblyShufps128Reverse
}

func (v *SHUFPS128REVERSE) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	shufps128Reverse(&vals1, &vals2, &ret)

	log.Printf("SHUFPS128REVERSE input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SHUFPS128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
