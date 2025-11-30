package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_shufps_128_twos.s
var assemblyShufps128Twos string

//go:embed stub_shufps_128_twos.go
var stubShufps128Twos string

type SHUFPS128TWOS struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewSHUFPS128TWOS() *SHUFPS128TWOS {
	return &SHUFPS128TWOS{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *SHUFPS128TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *SHUFPS128TWOS) Output() *number.Parameter {
	return v.ret
}

func (v *SHUFPS128TWOS) Name() string {
	return "SHUFPS (128 bit) twos"
}

func (v *SHUFPS128TWOS) Description() string {
	return "SHUFPS imm8=0xAA: dst = [a2,a2,b2,b2]"
}

func (v *SHUFPS128TWOS) Stub() string {
	return stubShufps128Twos
}

func (v *SHUFPS128TWOS) Assembly() string {
	return assemblyShufps128Twos
}

func (v *SHUFPS128TWOS) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	shufps128Twos(&vals1, &vals2, &ret)

	log.Printf("SHUFPS128TWOS input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *SHUFPS128TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
