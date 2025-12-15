package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_256_eight.s
var assemblyVpsrldq256Eight string

//go:embed stub_vpsrldq_256_eight.go
var stubVpsrldq256Eight string

type VPSRLDQ256EIGHT struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ256EIGHT() *VPSRLDQ256EIGHT {
	return &VPSRLDQ256EIGHT{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSRLDQ256EIGHT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ256EIGHT) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ256EIGHT) Name() string {
	return "VPSRLDQ (256 bit) eight"
}

func (v *VPSRLDQ256EIGHT) Description() string {
	return "Shift right by 8 bytes per 128-bit lane; upper 8 bytes of each lane become zero."
}

func (v *VPSRLDQ256EIGHT) Stub() string {
	return stubVpsrldq256Eight
}

func (v *VPSRLDQ256EIGHT) Assembly() string {
	return assemblyVpsrldq256Eight
}

func (v *VPSRLDQ256EIGHT) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq256Eight(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ256EIGHT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
