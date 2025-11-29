package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_128_mixed.s
var assemblyVshufps128Mixed string

//go:embed stub_vshufps_128_mixed.go
var stubVshufps128Mixed string

type VSHUFPS128MIXED struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS128MIXED() *VSHUFPS128MIXED {
	return &VSHUFPS128MIXED{
		vals1: number.NewNamedFloatParameter("vals1", 128, 32),
		vals2: number.NewNamedFloatParameter("vals2", 128, 32),
		ret:   number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VSHUFPS128MIXED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS128MIXED) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS128MIXED) Name() string {
	return "VSHUFPS (128 bit) mixed"
}

func (v *VSHUFPS128MIXED) Description() string {
	return "VSHUFPS imm8=0xE4: dst = [a0,a1,b2,b3]"
}

func (v *VSHUFPS128MIXED) Stub() string {
	return stubVshufps128Mixed
}

func (v *VSHUFPS128MIXED) Assembly() string {
	return assemblyVshufps128Mixed
}

func (v *VSHUFPS128MIXED) Run(_ [][]byte) (output []byte) {
	vals1 := [4]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [4]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [4]float32{}

	vshufps128Mixed(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS128MIXED input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS128MIXED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
