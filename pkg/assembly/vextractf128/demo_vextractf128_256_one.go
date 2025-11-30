package vextractf128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vextractf128_256_one.s
var assemblyVextractf128256One string

//go:embed stub_vextractf128_256_one.go
var stubVextractf128256One string

type VEXTRACTF128256ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVEXTRACTF128256ONE() *VEXTRACTF128256ONE {
	return &VEXTRACTF128256ONE{
		vals: number.NewNamedFloatParameter("vals", 256, 32),
		ret:  number.NewNamedFloatParameter("ret", 128, 32),
	}
}

func (v *VEXTRACTF128256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VEXTRACTF128256ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VEXTRACTF128256ONE) Name() string {
	return "VEXTRACTF128 (256 bit) one"
}

func (v *VEXTRACTF128256ONE) Description() string {
	return "Extract upper 128-bit lane (imm8=1) from YMM to XMM."
}

func (v *VEXTRACTF128256ONE) Stub() string {
	return stubVextractf128256One
}

func (v *VEXTRACTF128256ONE) Assembly() string {
	return assemblyVextractf128256One
}

func (v *VEXTRACTF128256ONE) Run() (output []byte) {
	base := [8]float32{}
	copy(base[:], number.ToFloat32Slice(v.vals.FlatData()))

	ret := [4]float32{}

	vextractf128256One(&base, &ret)

	log.Printf("VEXTRACTF128256ONE base %v output %v", base, ret)

	out := number.Float32SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VEXTRACTF128256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
