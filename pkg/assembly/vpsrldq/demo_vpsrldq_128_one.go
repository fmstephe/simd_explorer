package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_128_one.s
var assemblyVpsrldq128One string

//go:embed stub_vpsrldq_128_one.go
var stubVpsrldq128One string

type VPSRLDQ128ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ128ONE() *VPSRLDQ128ONE {
	return &VPSRLDQ128ONE{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSRLDQ128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ128ONE) Name() string {
	return "VPSRLDQ (128 bit) one"
}

func (v *VPSRLDQ128ONE) Description() string {
	return "Shift right by 1 byte within the 128-bit lane; highest byte becomes zero."
}

func (v *VPSRLDQ128ONE) Stub() string {
	return stubVpsrldq128One
}

func (v *VPSRLDQ128ONE) Assembly() string {
	return assemblyVpsrldq128One
}

func (v *VPSRLDQ128ONE) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq128One(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
