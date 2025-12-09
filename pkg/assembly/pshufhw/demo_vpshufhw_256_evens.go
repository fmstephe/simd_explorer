package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_256_evens.s
var assemblyVpshufhw256Evens string

//go:embed stub_vpshufhw_256_evens.go
var stubVpshufhw256Evens string

type VPSHUFHW256EVENS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW256EVENS() *VPSHUFHW256EVENS {
	return &VPSHUFHW256EVENS{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFHW256EVENS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW256EVENS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW256EVENS) Name() string {
	return "VPSHUFHW (256 bit) evens"
}

func (v *VPSHUFHW256EVENS) Description() string {
	return "Shuffle high words per 128-bit lane using imm8=0x88 (evens: [w4,w6,w4,w6]); low words unchanged."
}

func (v *VPSHUFHW256EVENS) Stub() string {
	return stubVpshufhw256Evens
}

func (v *VPSHUFHW256EVENS) Assembly() string {
	return assemblyVpshufhw256Evens
}

func (v *VPSHUFHW256EVENS) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshufhw256Evens(&vals, &ret)

	log.Printf("VPSHUFHW256EVENS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW256EVENS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
