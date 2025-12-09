package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_128_odds.s
var assemblyVpshuflw128Odds string

//go:embed stub_vpshuflw_128_odds.go
var stubVpshuflw128Odds string

type VPSHUFLW128ODDS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW128ODDS() *VPSHUFLW128ODDS {
	return &VPSHUFLW128ODDS{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFLW128ODDS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW128ODDS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW128ODDS) Name() string {
	return "VPSHUFLW (128 bit) odds"
}

func (v *VPSHUFLW128ODDS) Description() string {
	return "Shuffle low words in 128-bit lane using imm8=0xDD (odds: [w5,w7,w5,w7]); high words unchanged."
}

func (v *VPSHUFLW128ODDS) Stub() string {
	return stubVpshuflw128Odds
}

func (v *VPSHUFLW128ODDS) Assembly() string {
	return assemblyVpshuflw128Odds
}

func (v *VPSHUFLW128ODDS) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshuflw128Odds(&vals, &ret)

	log.Printf("VPSHUFLW128ODDS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW128ODDS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
