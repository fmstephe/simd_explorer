package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_128_odds.s
var assemblyVpshufd128Odds string

//go:embed stub_vpshufd_128_odds.go
var stubVpshufd128Odds string

type VPSHUFD128ODDS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD128ODDS() *VPSHUFD128ODDS {
	return &VPSHUFD128ODDS{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSHUFD128ODDS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD128ODDS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD128ODDS) Name() string {
	return "VPSHUFD ODDS (128 bit) "
}

func (v *VPSHUFD128ODDS) Description() string {
	return "Shuffle 32-bit integers within a 128-bit lane using imm8=0xDD (odds: [1,3,1,3])."
}

func (v *VPSHUFD128ODDS) Stub() string {
	return stubVpshufd128Odds
}

func (v *VPSHUFD128ODDS) Assembly() string {
	return assemblyVpshufd128Odds
}

func (v *VPSHUFD128ODDS) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [4]uint32{}

	vpshufd128Odds(&vals, &ret)

	log.Printf("VPSHUFD128ODDS vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD128ODDS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
