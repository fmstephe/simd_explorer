package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_256_sixteen.s
var assemblyVpsrldq256Sixteen string

//go:embed stub_vpsrldq_256_sixteen.go
var stubVpsrldq256Sixteen string

type VPSRLDQ256SIXTEEN struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ256SIXTEEN() *VPSRLDQ256SIXTEEN {
	return &VPSRLDQ256SIXTEEN{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSRLDQ256SIXTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ256SIXTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ256SIXTEEN) Name() string {
	return "VPSRLDQ (256 bit) sixteen"
}

func (v *VPSRLDQ256SIXTEEN) Description() string {
	return "Shift right by 16 bytes per 128-bit lane (full-lane zeroing)."
}

func (v *VPSRLDQ256SIXTEEN) Stub() string {
	return stubVpsrldq256Sixteen
}

func (v *VPSRLDQ256SIXTEEN) Assembly() string {
	return assemblyVpsrldq256Sixteen
}

func (v *VPSRLDQ256SIXTEEN) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq256Sixteen(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ256SIXTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
