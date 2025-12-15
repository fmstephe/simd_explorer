package vpsrldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpsrldq_256_one.s
var assemblyVpsrldq256One string

//go:embed stub_vpsrldq_256_one.go
var stubVpsrldq256One string

type VPSRLDQ256ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSRLDQ256ONE() *VPSRLDQ256ONE {
	return &VPSRLDQ256ONE{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSRLDQ256ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSRLDQ256ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSRLDQ256ONE) Name() string {
	return "VPSRLDQ (256 bit) one"
}

func (v *VPSRLDQ256ONE) Description() string {
	return "Shift right by 1 byte per 128-bit lane; highest byte of each lane becomes zero."
}

func (v *VPSRLDQ256ONE) Stub() string {
	return stubVpsrldq256One
}

func (v *VPSRLDQ256ONE) Assembly() string {
	return assemblyVpsrldq256One
}

func (v *VPSRLDQ256ONE) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpsrldq256One(&vals, &ret)

	log.Printf("VPSRLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSRLDQ256ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
