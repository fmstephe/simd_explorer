package shufps

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vshufps_256_zeros.s
var assemblyVshufps256Zeros string

//go:embed stub_vshufps_256_zeros.go
var stubVshufps256Zeros string

type VSHUFPS256ZEROS struct {
	vals1 *number.Parameter
	vals2 *number.Parameter
	ret   *number.Parameter
}

func NewVSHUFPS256ZEROS() *VSHUFPS256ZEROS {
	return &VSHUFPS256ZEROS{
		vals1: number.NewNamedFloatParameter("vals1", 256, 32),
		vals2: number.NewNamedFloatParameter("vals2", 256, 32),
		ret:   number.NewNamedFloatParameter("ret", 256, 32),
	}
}

func (v *VSHUFPS256ZEROS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals1,
		v.vals2,
	}
}

func (v *VSHUFPS256ZEROS) Output() *number.Parameter {
	return v.ret
}

func (v *VSHUFPS256ZEROS) Name() string {
	return "VSHUFPS (256 bit) zeros"
}

func (v *VSHUFPS256ZEROS) Description() string {
	return "VSHUFPS imm8=0x00 per 128-bit lane: dst = [a0,a0,b0,b0 | a4,a4,b4,b4]"
}

func (v *VSHUFPS256ZEROS) Stub() string {
	return stubVshufps256Zeros
}

func (v *VSHUFPS256ZEROS) Assembly() string {
	return assemblyVshufps256Zeros
}

func (v *VSHUFPS256ZEROS) Run(_ [][]byte) (output []byte) {
	vals1 := [8]float32{}
	copy(vals1[:], number.ToFloat32Slice(v.vals1.FlatData()))
	vals2 := [8]float32{}
	copy(vals2[:], number.ToFloat32Slice(v.vals2.FlatData()))

	ret := [8]float32{}

	vshufps256Zeros(&vals1, &vals2, &ret)

	log.Printf("VSHUFPS256ZEROS input %v %v output %v", vals1, vals2, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VSHUFPS256ZEROS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
