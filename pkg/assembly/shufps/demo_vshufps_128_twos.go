package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_twos.s
var assemblyVshufps128Twos string

//go:embed stub_vshufps_128_twos.go
var stubVshufps128Twos string

type VSHUFPS128TWOS struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS128TWOS() *VSHUFPS128TWOS {
	return &VSHUFPS128TWOS{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSHUFPS128TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS128TWOS) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS128TWOS) Name() string {
	return "VSHUFPS (128 bit) twos"
}

func (v *VSHUFPS128TWOS) Description() string {
	return "VSHUFPS imm8=0xAA: dst = [a2,a2,b2,b2]"
}

func (v *VSHUFPS128TWOS) Stub() string {
	return stubVshufps128Twos
}

func (v *VSHUFPS128TWOS) Assembly() string {
	return assemblyVshufps128Twos
}

func (v *VSHUFPS128TWOS) Run() {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vshufps128Twos(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS128TWOS input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VSHUFPS128TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
