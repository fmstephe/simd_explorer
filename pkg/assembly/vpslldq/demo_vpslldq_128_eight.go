package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_128_eight.s
var assemblyVpslldq128Eight string

//go:embed stub_vpslldq_128_eight.go
var stubVpslldq128Eight string

type VPSLLDQ128EIGHT struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ128EIGHT() *VPSLLDQ128EIGHT {
	return &VPSLLDQ128EIGHT{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSLLDQ128EIGHT) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ128EIGHT) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ128EIGHT) Name() string {
	return "VPSLLDQ (128 bit) eight"
}

func (v *VPSLLDQ128EIGHT) Description() string {
	return "Shift left by 8 bytes within the 128-bit lane; lower 8 bytes become zero."
}

func (v *VPSLLDQ128EIGHT) Stub() string {
	return stubVpslldq128Eight
}

func (v *VPSLLDQ128EIGHT) Assembly() string {
	return assemblyVpslldq128Eight
}

func (v *VPSLLDQ128EIGHT) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq128Eight(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ128EIGHT) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
