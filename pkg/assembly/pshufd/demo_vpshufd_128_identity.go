package pshufd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshufd_128_identity.s
var assemblyVpshufd128Identity string

//go:embed stub_vpshufd_128_identity.go
var stubVpshufd128Identity string

type VPSHUFD128IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFD128IDENTITY() *VPSHUFD128IDENTITY {
	return &VPSHUFD128IDENTITY{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPSHUFD128IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFD128IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFD128IDENTITY) Name() string {
	return "VPSHUFD IDENTITY (128 bit) "
}

func (v *VPSHUFD128IDENTITY) Description() string {
	return "Shuffle 32-bit integers within a 128-bit lane using imm8=0xE4 (identity)."
}

func (v *VPSHUFD128IDENTITY) Stub() string {
	return stubVpshufd128Identity
}

func (v *VPSHUFD128IDENTITY) Assembly() string {
	return assemblyVpshufd128Identity
}

func (v *VPSHUFD128IDENTITY) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))

	ret := [4]uint32{}

	vpshufd128Identity(&vals, &ret)

	log.Printf("VPSHUFD128IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFD128IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
