package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_reverse.s
var assemblyVshufps128Reverse string

//go:embed stub_vshufps_128_reverse.go
var stubVshufps128Reverse string

type VSHUFPS128REVERSE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS128REVERSE() *VSHUFPS128REVERSE {
	return &VSHUFPS128REVERSE{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSHUFPS128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS128REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS128REVERSE) Name() string {
	return "VSHUFPS (128 bit) reverse"
}

func (v *VSHUFPS128REVERSE) Description() string {
	return "VSHUFPS imm8=0x1B: dst = [a3,a2,b1,b0]"
}

func (v *VSHUFPS128REVERSE) Stub() string {
	return stubVshufps128Reverse
}

func (v *VSHUFPS128REVERSE) Assembly() string {
	return assemblyVshufps128Reverse
}

func (v *VSHUFPS128REVERSE) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vshufps128Reverse(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS128REVERSE input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
