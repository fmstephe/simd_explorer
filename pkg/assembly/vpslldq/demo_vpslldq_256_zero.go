package vpslldq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpslldq_256_zero.s
var assemblyVpslldq256Zero string

//go:embed stub_vpslldq_256_zero.go
var stubVpslldq256Zero string

type VPSLLDQ256ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSLLDQ256ZERO() *VPSLLDQ256ZERO {
	return &VPSLLDQ256ZERO{
		vals: number.NewNamedUintParameter("vals", 256, 8, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 8, 10),
	}
}

func (v *VPSLLDQ256ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSLLDQ256ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPSLLDQ256ZERO) Name() string {
	return "VPSLLDQ (256 bit) zero"
}

func (v *VPSLLDQ256ZERO) Description() string {
	return "Shift left by 0 bytes per 128-bit lane (no change)."
}

func (v *VPSLLDQ256ZERO) Stub() string {
	return stubVpslldq256Zero
}

func (v *VPSLLDQ256ZERO) Assembly() string {
	return assemblyVpslldq256Zero
}

func (v *VPSLLDQ256ZERO) Run() {
	vals := [32]uint8{}
	copy(vals[:], number.ToUint8Slice(v.vals.FlatData()))
	ret := [32]uint8{}
	copy(ret[:], number.ToUint8Slice(v.ret.FlatData()))

	vpslldq256Zero(&vals, &ret)

	log.Printf("VPSLLDQ vals %v ret %v", vals, ret)

	retBytes := number.Uint8SliceToBytes(ret[:])
	v.ret.SetData(retBytes)
}

func (v *VPSLLDQ256ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
