package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_128_fifteen.s
var assemblyVpsrldq128Fifteen string

//go:embed stub_vpsrldq_128_fifteen.go
var stubVpsrldq128Fifteen string

type VPSRLDQ128FIFTEEN struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ128FIFTEEN() *VPSRLDQ128FIFTEEN {
	return &VPSRLDQ128FIFTEEN{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSRLDQ128FIFTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ128FIFTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ128FIFTEEN) Name() string {
	return "VPSRLDQ (128 bit) fifteen"
}

func (v *VPSRLDQ128FIFTEEN) Description() string {
	return "Shift right by 15 bytes within the 128-bit lane; only the lowest byte may survive."
}

func (v *VPSRLDQ128FIFTEEN) Stub() string {
	return stubVpsrldq128Fifteen
}

func (v *VPSRLDQ128FIFTEEN) Assembly() string {
	return assemblyVpsrldq128Fifteen
}

func (v *VPSRLDQ128FIFTEEN) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq128Fifteen(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ128FIFTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
