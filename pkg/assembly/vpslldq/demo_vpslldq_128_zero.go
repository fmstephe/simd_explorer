package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_128_zero.s
var assemblyVpslldq128Zero string

//go:embed stub_vpslldq_128_zero.go
var stubVpslldq128Zero string

type VPSLLDQ128ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ128ZERO() *VPSLLDQ128ZERO {
	return &VPSLLDQ128ZERO{
		vals: number.NewNamedUintParameter("vals", 128, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 8, 10),
	}
}

func (v *VPSLLDQ128ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ128ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ128ZERO) Name() string {
	return "VPSLLDQ (128 bit) zero"
}

func (v *VPSLLDQ128ZERO) Description() string {
	return "Shift left by 0 bytes within the 128-bit lane (no change)."
}

func (v *VPSLLDQ128ZERO) Stub() string {
	return stubVpslldq128Zero
}

func (v *VPSLLDQ128ZERO) Assembly() string {
	return assemblyVpslldq128Zero
}

func (v *VPSLLDQ128ZERO) Run() {
	vals := [16]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [16]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq128Zero(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ128ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
