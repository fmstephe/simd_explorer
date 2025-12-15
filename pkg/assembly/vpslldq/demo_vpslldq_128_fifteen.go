package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_128_fifteen.s
var assemblyVpslldq128Fifteen string

//go:embed stub_vpslldq_128_fifteen.go
var stubVpslldq128Fifteen string

type VPSLLDQ128FIFTEEN struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ128FIFTEEN() *VPSLLDQ128FIFTEEN {
	return &VPSLLDQ128FIFTEEN{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSLLDQ128FIFTEEN) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ128FIFTEEN) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ128FIFTEEN) Name() string {
	return "VPSLLDQ (128 bit) fifteen"
}

func (v *VPSLLDQ128FIFTEEN) Description() string {
	return "Shift left by 15 bytes within the 128-bit lane; only the highest byte may survive."
}

func (v *VPSLLDQ128FIFTEEN) Stub() string {
	return stubVpslldq128Fifteen
}

func (v *VPSLLDQ128FIFTEEN) Assembly() string {
	return assemblyVpslldq128Fifteen
}

func (v *VPSLLDQ128FIFTEEN) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq128Fifteen(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ128FIFTEEN) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
