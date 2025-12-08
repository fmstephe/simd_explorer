package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_256_odds.s
var assemblyVpshufd256Odds string

//go:embed stub_vpshufd_256_odds.go
var stubVpshufd256Odds string

type VPSHUFD256ODDS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD256ODDS() *VPSHUFD256ODDS {
	return &VPSHUFD256ODDS{
		vals: number.NewNamedUintParameter("vals", 256, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSHUFD256ODDS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD256ODDS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD256ODDS) Name() string {
	return "VPSHUFD ODDS (256 bit) "
}

func (v *VPSHUFD256ODDS) Description() string {
	return "Shuffle 32-bit integers within each 128-bit lane using imm8=0xDD (odds: [1,3,1,3])."
}

func (v *VPSHUFD256ODDS) Stub() string {
	return stubVpshufd256Odds
}

func (v *VPSHUFD256ODDS) Assembly() string {
	return assemblyVpshufd256Odds
}

func (v *VPSHUFD256ODDS) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [8]uint32{}

	vpshufd256Odds(&vals, &ret)

	log.Printf("VPSHUFD256ODDS vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD256ODDS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
