package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_128_identity.s
var assemblyVpshufhw128Identity string

//go:embed stub_vpshufhw_128_identity.go
var stubVpshufhw128Identity string

type VPSHUFHW128IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW128IDENTITY() *VPSHUFHW128IDENTITY {
	return &VPSHUFHW128IDENTITY{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFHW128IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW128IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW128IDENTITY) Name() string {
	return "VPSHUFHW (128 bit) identity"
}

func (v *VPSHUFHW128IDENTITY) Description() string {
	return "Shuffle high words in 128-bit lane using imm8=0xE4 (identity); low words unchanged."
}

func (v *VPSHUFHW128IDENTITY) Stub() string {
	return stubVpshufhw128Identity
}

func (v *VPSHUFHW128IDENTITY) Assembly() string {
	return assemblyVpshufhw128Identity
}

func (v *VPSHUFHW128IDENTITY) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshufhw128Identity(&vals, &ret)

	log.Printf("VPSHUFHW128IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW128IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
