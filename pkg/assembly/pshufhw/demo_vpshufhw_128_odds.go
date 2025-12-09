package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_128_odds.s
var assemblyVpshufhw128Odds string

//go:embed stub_vpshufhw_128_odds.go
var stubVpshufhw128Odds string

type VPSHUFHW128ODDS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW128ODDS() *VPSHUFHW128ODDS {
	return &VPSHUFHW128ODDS{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFHW128ODDS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW128ODDS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW128ODDS) Name() string {
	return "VPSHUFHW (128 bit) odds"
}

func (v *VPSHUFHW128ODDS) Description() string {
	return "Shuffle high words in 128-bit lane using imm8=0xDD (odds: [w5,w7,w5,w7]); low words unchanged."
}

func (v *VPSHUFHW128ODDS) Stub() string {
	return stubVpshufhw128Odds
}

func (v *VPSHUFHW128ODDS) Assembly() string {
	return assemblyVpshufhw128Odds
}

func (v *VPSHUFHW128ODDS) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshufhw128Odds(&vals, &ret)

	log.Printf("VPSHUFHW128ODDS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW128ODDS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
