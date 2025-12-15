package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_256_fifteen.s
var assemblyVpsrldq256Fifteen string

//go:embed stub_vpsrldq_256_fifteen.go
var stubVpsrldq256Fifteen string

type VPSRLDQ256FIFTEEN struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ256FIFTEEN() *VPSRLDQ256FIFTEEN {
	return &VPSRLDQ256FIFTEEN{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSRLDQ256FIFTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ256FIFTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ256FIFTEEN) Name() string {
	return "VPSRLDQ (256 bit) fifteen"
}

func (v *VPSRLDQ256FIFTEEN) Description() string {
	return "Shift right by 15 bytes per 128-bit lane; only lowest byte of each lane may survive."
}

func (v *VPSRLDQ256FIFTEEN) Stub() string {
	return stubVpsrldq256Fifteen
}

func (v *VPSRLDQ256FIFTEEN) Assembly() string {
	return assemblyVpsrldq256Fifteen
}

func (v *VPSRLDQ256FIFTEEN) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq256Fifteen(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ256FIFTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
