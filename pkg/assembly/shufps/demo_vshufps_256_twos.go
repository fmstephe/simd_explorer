package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_twos.s
var assemblyVshufps256Twos string

//go:embed stub_vshufps_256_twos.go
var stubVshufps256Twos string

type VSHUFPS256TWOS struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS256TWOS() *VSHUFPS256TWOS {
	return &VSHUFPS256TWOS{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSHUFPS256TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS256TWOS) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS256TWOS) Name() string {
	return "VSHUFPS (256 bit) twos"
}

func (v *VSHUFPS256TWOS) Description() string {
	return "VSHUFPS imm8=0xAA per 128-bit lane: dst = [a2,a2,b2,b2 | a6,a6,b6,b6]"
}

func (v *VSHUFPS256TWOS) Stub() string {
	return stubVshufps256Twos
}

func (v *VSHUFPS256TWOS) Assembly() string {
	return assemblyVshufps256Twos
}

func (v *VSHUFPS256TWOS) Run() (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vshufps256Twos(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS256TWOS input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS256TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
