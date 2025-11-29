package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_zeros.s
var assemblyShufps128Zeros string

//go:embed stub_shufps_128_zeros.go
var stubShufps128Zeros string

type SHUFPS128ZEROS struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSHUFPS128ZEROS() *SHUFPS128ZEROS {
	return &SHUFPS128ZEROS{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SHUFPS128ZEROS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SHUFPS128ZEROS) Output() *number.Parameter {
	return v.ret
}

func (v *SHUFPS128ZEROS) Name() string {
	return "SHUFPS (128 bit) zeros"
}

func (v *SHUFPS128ZEROS) Description() string {
	return "SHUFPS imm8=0x00: dst = [a0,a0,b0,b0]"
}

func (v *SHUFPS128ZEROS) Stub() string {
	return stubShufps128Zeros
}

func (v *SHUFPS128ZEROS) Assembly() string {
	return assemblyShufps128Zeros
}

func (v *SHUFPS128ZEROS) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	shufps128Zeros(&vals1, &vals2, &ret)

	log.Printf("SHUFPS128ZEROS input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *SHUFPS128ZEROS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
