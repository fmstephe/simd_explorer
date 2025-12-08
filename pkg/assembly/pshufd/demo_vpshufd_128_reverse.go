package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_128_reverse.s
var assemblyVpshufd128Reverse string

//go:embed stub_vpshufd_128_reverse.go
var stubVpshufd128Reverse string

type VPSHUFD128REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD128REVERSE() *VPSHUFD128REVERSE {
	return &VPSHUFD128REVERSE{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSHUFD128REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD128REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD128REVERSE) Name() string {
	return "VPSHUFD REVERSE (128 bit) "
}

func (v *VPSHUFD128REVERSE) Description() string {
	return "Shuffle 32-bit integers within a 128-bit lane using imm8=0x1B (reverse)."
}

func (v *VPSHUFD128REVERSE) Stub() string {
	return stubVpshufd128Reverse
}

func (v *VPSHUFD128REVERSE) Assembly() string {
	return assemblyVpshufd128Reverse
}

func (v *VPSHUFD128REVERSE) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [4]uint32{}

	vpshufd128Reverse(&vals, &ret)

	log.Printf("VPSHUFD128REVERSE vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD128REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
