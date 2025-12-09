package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_256_reverse.s
var assemblyVpshuflw256Reverse string

//go:embed stub_vpshuflw_256_reverse.go
var stubVpshuflw256Reverse string

type VPSHUFLW256REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW256REVERSE() *VPSHUFLW256REVERSE {
	return &VPSHUFLW256REVERSE{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFLW256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW256REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW256REVERSE) Name() string {
	return "VPSHUFLW (256 bit) reverse"
}

func (v *VPSHUFLW256REVERSE) Description() string {
	return "Shuffle low words per 128-bit lane using imm8=0x1B (reverse); high words unchanged."
}

func (v *VPSHUFLW256REVERSE) Stub() string {
	return stubVpshuflw256Reverse
}

func (v *VPSHUFLW256REVERSE) Assembly() string {
	return assemblyVpshuflw256Reverse
}

func (v *VPSHUFLW256REVERSE) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshuflw256Reverse(&vals, &ret)

	log.Printf("VPSHUFLW256REVERSE vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
