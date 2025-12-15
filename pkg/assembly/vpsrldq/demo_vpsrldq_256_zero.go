package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_256_zero.s
var assemblyVpsrldq256Zero string

//go:embed stub_vpsrldq_256_zero.go
var stubVpsrldq256Zero string

type VPSRLDQ256ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ256ZERO() *VPSRLDQ256ZERO {
	return &VPSRLDQ256ZERO{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSRLDQ256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ256ZERO) Name() string {
	return "VPSRLDQ (256 bit) zero"
}

func (v *VPSRLDQ256ZERO) Description() string {
	return "Shift right by 0 bytes per 128-bit lane (no change)."
}

func (v *VPSRLDQ256ZERO) Stub() string {
	return stubVpsrldq256Zero
}

func (v *VPSRLDQ256ZERO) Assembly() string {
	return assemblyVpsrldq256Zero
}

func (v *VPSRLDQ256ZERO) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq256Zero(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
