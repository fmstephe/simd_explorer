package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_128_zero.s
var assemblyVpsrldq128Zero string

//go:embed stub_vpsrldq_128_zero.go
var stubVpsrldq128Zero string

type VPSRLDQ128ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ128ZERO() *VPSRLDQ128ZERO {
	return &VPSRLDQ128ZERO{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSRLDQ128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ128ZERO) Name() string {
	return "VPSRLDQ (128 bit) zero"
}

func (v *VPSRLDQ128ZERO) Description() string {
	return "Shift right by 0 bytes within the 128-bit lane (no change)."
}

func (v *VPSRLDQ128ZERO) Stub() string {
	return stubVpsrldq128Zero
}

func (v *VPSRLDQ128ZERO) Assembly() string {
	return assemblyVpsrldq128Zero
}

func (v *VPSRLDQ128ZERO) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq128Zero(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
