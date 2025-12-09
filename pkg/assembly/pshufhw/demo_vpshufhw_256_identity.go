package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_256_identity.s
var assemblyVpshufhw256Identity string

//go:embed stub_vpshufhw_256_identity.go
var stubVpshufhw256Identity string

type VPSHUFHW256IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW256IDENTITY() *VPSHUFHW256IDENTITY {
	return &VPSHUFHW256IDENTITY{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFHW256IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW256IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW256IDENTITY) Name() string {
	return "VPSHUFHW (256 bit) identity"
}

func (v *VPSHUFHW256IDENTITY) Description() string {
	return "Shuffle high words per 128-bit lane using imm8=0xE4 (identity); low words unchanged."
}

func (v *VPSHUFHW256IDENTITY) Stub() string {
	return stubVpshufhw256Identity
}

func (v *VPSHUFHW256IDENTITY) Assembly() string {
	return assemblyVpshufhw256Identity
}

func (v *VPSHUFHW256IDENTITY) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshufhw256Identity(&vals, &ret)

	log.Printf("VPSHUFHW256IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW256IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
