package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_128_evens.s
var assemblyVpshufhw128Evens string

//go:embed stub_vpshufhw_128_evens.go
var stubVpshufhw128Evens string

type VPSHUFHW128EVENS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW128EVENS() *VPSHUFHW128EVENS {
	return &VPSHUFHW128EVENS{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFHW128EVENS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW128EVENS) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW128EVENS) Name() string {
	return "VPSHUFHW (128 bit) evens"
}

func (v *VPSHUFHW128EVENS) Description() string {
	return "Shuffle high words in 128-bit lane using imm8=0x88 (evens: [w4,w6,w4,w6]); low words unchanged."
}

func (v *VPSHUFHW128EVENS) Stub() string {
	return stubVpshufhw128Evens
}

func (v *VPSHUFHW128EVENS) Assembly() string {
	return assemblyVpshufhw128Evens
}

func (v *VPSHUFHW128EVENS) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshufhw128Evens(&vals, &ret)

	log.Printf("VPSHUFHW128EVENS vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW128EVENS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
