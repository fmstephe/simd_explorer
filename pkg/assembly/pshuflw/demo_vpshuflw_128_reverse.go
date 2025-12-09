package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_128_reverse.s
var assemblyVpshuflw128Reverse string

//go:embed stub_vpshuflw_128_reverse.go
var stubVpshuflw128Reverse string

type VPSHUFLW128REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW128REVERSE() *VPSHUFLW128REVERSE {
	return &VPSHUFLW128REVERSE{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFLW128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW128REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW128REVERSE) Name() string {
	return "VPSHUFLW (128 bit) reverse"
}

func (v *VPSHUFLW128REVERSE) Description() string {
	return "Shuffle low words in 128-bit lane using imm8=0x1B (reverse); high words unchanged."
}

func (v *VPSHUFLW128REVERSE) Stub() string {
	return stubVpshuflw128Reverse
}

func (v *VPSHUFLW128REVERSE) Assembly() string {
	return assemblyVpshuflw128Reverse
}

func (v *VPSHUFLW128REVERSE) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshuflw128Reverse(&vals, &ret)

	log.Printf("VPSHUFLW128REVERSE vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
