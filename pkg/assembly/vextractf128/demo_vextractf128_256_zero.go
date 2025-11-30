package vextractf128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vextractf128_256_zero.s
var assemblyVextractf128256Zero string

//go:embed stub_vextractf128_256_zero.go
var stubVextractf128256Zero string

type VEXTRACTF128256ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVEXTRACTF128256ZERO() *VEXTRACTF128256ZERO {
	return &VEXTRACTF128256ZERO{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VEXTRACTF128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VEXTRACTF128256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VEXTRACTF128256ZERO) Name() string {
	return "VEXTRACTF128 (256 bit) zero"
}

func (v *VEXTRACTF128256ZERO) Description() string {
	return "Extract lower 128-bit lane (imm8=0) from YMM to XMM."
}

func (v *VEXTRACTF128256ZERO) Stub() string {
	return stubVextractf128256Zero
}

func (v *VEXTRACTF128256ZERO) Assembly() string {
	return assemblyVextractf128256Zero
}

func (v *VEXTRACTF128256ZERO) Run() (output []byte) {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vextractf128256Zero(&base, &ret)

	log.Printf("VEXTRACTF128256ZERO base %v output %v", base, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VEXTRACTF128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
