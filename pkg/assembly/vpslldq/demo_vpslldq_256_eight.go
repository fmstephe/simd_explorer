package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_256_eight.s
var assemblyVpslldq256Eight string

//go:embed stub_vpslldq_256_eight.go
var stubVpslldq256Eight string

type VPSLLDQ256EIGHT struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ256EIGHT() *VPSLLDQ256EIGHT {
	return &VPSLLDQ256EIGHT{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSLLDQ256EIGHT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ256EIGHT) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ256EIGHT) Name() string {
	return "VPSLLDQ (256 bit) eight"
}

func (v *VPSLLDQ256EIGHT) Description() string {
	return "Shift left by 8 bytes per 128-bit lane; lower 8 bytes of each lane become zero."
}

func (v *VPSLLDQ256EIGHT) Stub() string {
	return stubVpslldq256Eight
}

func (v *VPSLLDQ256EIGHT) Assembly() string {
	return assemblyVpslldq256Eight
}

func (v *VPSLLDQ256EIGHT) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq256Eight(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ256EIGHT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
