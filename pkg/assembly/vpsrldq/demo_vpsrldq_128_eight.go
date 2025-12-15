package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_128_eight.s
var assemblyVpsrldq128Eight string

//go:embed stub_vpsrldq_128_eight.go
var stubVpsrldq128Eight string

type VPSRLDQ128EIGHT struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ128EIGHT() *VPSRLDQ128EIGHT {
	return &VPSRLDQ128EIGHT{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSRLDQ128EIGHT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ128EIGHT) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ128EIGHT) Name() string {
	return "VPSRLDQ (128 bit) eight"
}

func (v *VPSRLDQ128EIGHT) Description() string {
	return "Shift right by 8 bytes within the 128-bit lane; upper 8 bytes become zero."
}

func (v *VPSRLDQ128EIGHT) Stub() string {
	return stubVpsrldq128Eight
}

func (v *VPSRLDQ128EIGHT) Assembly() string {
	return assemblyVpsrldq128Eight
}

func (v *VPSRLDQ128EIGHT) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq128Eight(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ128EIGHT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
