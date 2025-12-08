package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_256_identity.s
var assemblyVpshufd256Identity string

//go:embed stub_vpshufd_256_identity.go
var stubVpshufd256Identity string

type VPSHUFD256IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD256IDENTITY() *VPSHUFD256IDENTITY {
	return &VPSHUFD256IDENTITY{
		vals: number.NewNamedUintParameter("vals", 256, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPSHUFD256IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD256IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD256IDENTITY) Name() string {
	return "VPSHUFD IDENTITY (256 bit) "
}

func (v *VPSHUFD256IDENTITY) Description() string {
	return "Shuffle 32-bit integers within each 128-bit lane using imm8=0xE4 (identity)."
}

func (v *VPSHUFD256IDENTITY) Stub() string {
	return stubVpshufd256Identity
}

func (v *VPSHUFD256IDENTITY) Assembly() string {
	return assemblyVpshufd256Identity
}

func (v *VPSHUFD256IDENTITY) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [8]uint32{}

	vpshufd256Identity(&vals, &ret)

	log.Printf("VPSHUFD256IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD256IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
