package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_mixed.s
var assemblyVshufps256Mixed string

//go:embed stub_vshufps_256_mixed.go
var stubVshufps256Mixed string

type VSHUFPS256MIXED struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS256MIXED() *VSHUFPS256MIXED {
	return &VSHUFPS256MIXED{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSHUFPS256MIXED) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS256MIXED) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS256MIXED) Name() string {
	return "VSHUFPS (256 bit) mixed"
}

func (v *VSHUFPS256MIXED) Description() string {
	return "VSHUFPS imm8=0xE4 per 128-bit lane: dst = [a0,a1,b2,b3 | a4,a5,b6,b7]"
}

func (v *VSHUFPS256MIXED) Stub() string {
	return stubVshufps256Mixed
}

func (v *VSHUFPS256MIXED) Assembly() string {
	return assemblyVshufps256Mixed
}

func (v *VSHUFPS256MIXED) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vshufps256Mixed(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS256MIXED input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS256MIXED) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
