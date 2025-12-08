package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_256_reverse.s
var assemblyVpshufd256Reverse string

//go:embed stub_vpshufd_256_reverse.go
var stubVpshufd256Reverse string

type VPSHUFD256REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD256REVERSE() *VPSHUFD256REVERSE {
	return &VPSHUFD256REVERSE{
		vals: number.NewNamedUintParameter("vals", 256, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSHUFD256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD256REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD256REVERSE) Name() string {
	return "VPSHUFD REVERSE (256 bit) "
}

func (v *VPSHUFD256REVERSE) Description() string {
	return "Shuffle 32-bit integers within each 128-bit lane using imm8=0x1B (reverse)."
}

func (v *VPSHUFD256REVERSE) Stub() string {
	return stubVpshufd256Reverse
}

func (v *VPSHUFD256REVERSE) Assembly() string {
	return assemblyVpshufd256Reverse
}

func (v *VPSHUFD256REVERSE) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [8]uint32{}

	vpshufd256Reverse(&vals, &ret)

	log.Printf("VPSHUFD256REVERSE vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
