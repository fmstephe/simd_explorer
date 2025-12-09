package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_256_odds.s
var assemblyVpshuflw256Odds string

//go:embed stub_vpshuflw_256_odds.go
var stubVpshuflw256Odds string

type VPSHUFLW256ODDS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW256ODDS() *VPSHUFLW256ODDS {
	return &VPSHUFLW256ODDS{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFLW256ODDS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW256ODDS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW256ODDS) Name() string {
	return "VPSHUFLW (256 bit) odds"
}

func (v *VPSHUFLW256ODDS) Description() string {
	return "Shuffle low words per 128-bit lane using imm8=0xDD (odds: [w5,w7,w5,w7]); high words unchanged."
}

func (v *VPSHUFLW256ODDS) Stub() string {
	return stubVpshuflw256Odds
}

func (v *VPSHUFLW256ODDS) Assembly() string {
	return assemblyVpshuflw256Odds
}

func (v *VPSHUFLW256ODDS) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshuflw256Odds(&vals, &ret)

	log.Printf("VPSHUFLW256ODDS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW256ODDS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
