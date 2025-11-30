package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_reverse.s
var assemblyVshufps256Reverse string

//go:embed stub_vshufps_256_reverse.go
var stubVshufps256Reverse string

type VSHUFPS256REVERSE struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS256REVERSE() *VSHUFPS256REVERSE {
	return &VSHUFPS256REVERSE{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSHUFPS256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS256REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS256REVERSE) Name() string {
	return "VSHUFPS (256 bit) reverse"
}

func (v *VSHUFPS256REVERSE) Description() string {
	return "VSHUFPS imm8=0x1B per 128-bit lane: dst = [a3,a2,b1,b0 | a7,a6,b5,b4]"
}

func (v *VSHUFPS256REVERSE) Stub() string {
	return stubVshufps256Reverse
}

func (v *VSHUFPS256REVERSE) Assembly() string {
	return assemblyVshufps256Reverse
}

func (v *VSHUFPS256REVERSE) Run() (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vshufps256Reverse(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS256REVERSE input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
