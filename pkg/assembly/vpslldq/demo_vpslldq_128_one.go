package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_128_one.s
var assemblyVpslldq128One string

//go:embed stub_vpslldq_128_one.go
var stubVpslldq128One string

type VPSLLDQ128ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ128ONE() *VPSLLDQ128ONE {
	return &VPSLLDQ128ONE{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSLLDQ128ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ128ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ128ONE) Name() string {
	return "VPSLLDQ (128 bit) one"
}

func (v *VPSLLDQ128ONE) Description() string {
	return "Shift left by 1 byte within the 128-bit lane; lowest byte becomes zero."
}

func (v *VPSLLDQ128ONE) Stub() string {
	return stubVpslldq128One
}

func (v *VPSLLDQ128ONE) Assembly() string {
	return assemblyVpslldq128One
}

func (v *VPSLLDQ128ONE) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq128One(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ128ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
