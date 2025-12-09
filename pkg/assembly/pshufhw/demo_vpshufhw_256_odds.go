package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_256_odds.s
var assemblyVpshufhw256Odds string

//go:embed stub_vpshufhw_256_odds.go
var stubVpshufhw256Odds string

type VPSHUFHW256ODDS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW256ODDS() *VPSHUFHW256ODDS {
	return &VPSHUFHW256ODDS{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFHW256ODDS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW256ODDS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW256ODDS) Name() string {
	return "VPSHUFHW (256 bit) odds"
}

func (v *VPSHUFHW256ODDS) Description() string {
	return "Shuffle high words per 128-bit lane using imm8=0xDD (odds: [w5,w7,w5,w7]); low words unchanged."
}

func (v *VPSHUFHW256ODDS) Stub() string {
	return stubVpshufhw256Odds
}

func (v *VPSHUFHW256ODDS) Assembly() string {
	return assemblyVpshufhw256Odds
}

func (v *VPSHUFHW256ODDS) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshufhw256Odds(&vals, &ret)

	log.Printf("VPSHUFHW256ODDS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW256ODDS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
