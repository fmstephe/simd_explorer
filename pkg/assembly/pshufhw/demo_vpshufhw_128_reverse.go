package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_128_reverse.s
var assemblyVpshufhw128Reverse string

//go:embed stub_vpshufhw_128_reverse.go
var stubVpshufhw128Reverse string

type VPSHUFHW128REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW128REVERSE() *VPSHUFHW128REVERSE {
	return &VPSHUFHW128REVERSE{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFHW128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW128REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW128REVERSE) Name() string {
	return "VPSHUFHW (128 bit) reverse"
}

func (v *VPSHUFHW128REVERSE) Description() string {
	return "Shuffle high words in 128-bit lane using imm8=0x1B (reverse); low words unchanged."
}

func (v *VPSHUFHW128REVERSE) Stub() string {
	return stubVpshufhw128Reverse
}

func (v *VPSHUFHW128REVERSE) Assembly() string {
	return assemblyVpshufhw128Reverse
}

func (v *VPSHUFHW128REVERSE) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshufhw128Reverse(&vals, &ret)

	log.Printf("VPSHUFHW128REVERSE vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
