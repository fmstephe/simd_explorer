package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_mixed.s
var assemblyShufps128Mixed string

//go:embed stub_shufps_128_mixed.go
var stubShufps128Mixed string

type SHUFPS128MIXED struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSHUFPS128MIXED() *SHUFPS128MIXED {
	return &SHUFPS128MIXED{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SHUFPS128MIXED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SHUFPS128MIXED) Output() *number.Parameter {
	return v.ret
}

func (v *SHUFPS128MIXED) Name() string {
	return "SHUFPS (128 bit) mixed"
}

func (v *SHUFPS128MIXED) Description() string {
	return "SHUFPS imm8=0xE4: dst = [a0,a1,b2,b3]"
}

func (v *SHUFPS128MIXED) Stub() string {
	return stubShufps128Mixed
}

func (v *SHUFPS128MIXED) Assembly() string {
	return assemblyShufps128Mixed
}

func (v *SHUFPS128MIXED) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	shufps128Mixed(&vals1, &vals2, &ret)

	log.Printf("SHUFPS128MIXED input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SHUFPS128MIXED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
