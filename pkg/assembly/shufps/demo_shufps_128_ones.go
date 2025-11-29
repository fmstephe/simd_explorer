package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_ones.s
var assemblyShufps128Ones string

//go:embed stub_shufps_128_ones.go
var stubShufps128Ones string

type SHUFPS128ONES struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSHUFPS128ONES() *SHUFPS128ONES {
	return &SHUFPS128ONES{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SHUFPS128ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SHUFPS128ONES) Output() *number.Parameter {
	return v.ret
}

func (v *SHUFPS128ONES) Name() string {
	return "SHUFPS (128 bit) ones"
}

func (v *SHUFPS128ONES) Description() string {
	return "SHUFPS imm8=0x55: dst = [a1,a1,b1,b1]"
}

func (v *SHUFPS128ONES) Stub() string {
	return stubShufps128Ones
}

func (v *SHUFPS128ONES) Assembly() string {
	return assemblyShufps128Ones
}

func (v *SHUFPS128ONES) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	shufps128Ones(&vals1, &vals2, &ret)

	log.Printf("SHUFPS128ONES input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SHUFPS128ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
