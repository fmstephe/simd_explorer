package vextracti128

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vextracti128_256_zero.s
var assemblyVextracti128256Zero string

//go:embed stub_vextracti128_256_zero.go
var stubVextracti128256Zero string

type VEXTRACTI128256ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVEXTRACTI128256ZERO() *VEXTRACTI128256ZERO {
	return &VEXTRACTI128256ZERO{
		vals: number.NewNamedUintParameter("vals", 256, 32, 16),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 16),
	}
}

func (v *VEXTRACTI128256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VEXTRACTI128256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VEXTRACTI128256ZERO) Name() string {
	return "VEXTRACTI128 (256 bit) zero"
}

func (v *VEXTRACTI128256ZERO) Description() string {
	return "Extract lower 128-bit lane (0) from YMM to memory."
}

func (v *VEXTRACTI128256ZERO) Stub() string {
	return stubVextracti128256Zero
}

func (v *VEXTRACTI128256ZERO) Assembly() string {
	return assemblyVextracti128256Zero
}

func (v *VEXTRACTI128256ZERO) Run() {
	var vals256 [8]uint32
	copy(vals256[:], number.ToUint32Slice(v.vals.FlatData()))
	var ret [4]uint32
	vextracti128256Zero(&vals256, &ret)
	log.Printf("VEXTRACTI128256ZERO vals256 %v output %v", vals256, ret)
	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VEXTRACTI128256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
