package pshufhw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufhw_256_reverse.s
var assemblyVpshufhw256Reverse string

//go:embed stub_vpshufhw_256_reverse.go
var stubVpshufhw256Reverse string

type VPSHUFHW256REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFHW256REVERSE() *VPSHUFHW256REVERSE {
	return &VPSHUFHW256REVERSE{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFHW256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFHW256REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFHW256REVERSE) Name() string {
	return "VPSHUFHW (256 bit) reverse"
}

func (v *VPSHUFHW256REVERSE) Description() string {
	return "Shuffle high words per 128-bit lane using imm8=0x1B (reverse); low words unchanged."
}

func (v *VPSHUFHW256REVERSE) Stub() string {
	return stubVpshufhw256Reverse
}

func (v *VPSHUFHW256REVERSE) Assembly() string {
	return assemblyVpshufhw256Reverse
}

func (v *VPSHUFHW256REVERSE) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshufhw256Reverse(&vals, &ret)

	log.Printf("VPSHUFHW256REVERSE vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFHW256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
