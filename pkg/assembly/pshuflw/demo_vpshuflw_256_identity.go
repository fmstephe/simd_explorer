package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_256_identity.s
var assemblyVpshuflw256Identity string

//go:embed stub_vpshuflw_256_identity.go
var stubVpshuflw256Identity string

type VPSHUFLW256IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW256IDENTITY() *VPSHUFLW256IDENTITY {
	return &VPSHUFLW256IDENTITY{
		vals: number.NewNamedUintParameter("vals", 256, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 16, 10),
	}
}

func (v *VPSHUFLW256IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW256IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW256IDENTITY) Name() string {
	return "VPSHUFLW (256 bit) identity"
}

func (v *VPSHUFLW256IDENTITY) Description() string {
	return "Shuffle low words per 128-bit lane using imm8=0xE4 (identity); high words unchanged."
}

func (v *VPSHUFLW256IDENTITY) Stub() string {
	return stubVpshuflw256Identity
}

func (v *VPSHUFLW256IDENTITY) Assembly() string {
	return assemblyVpshuflw256Identity
}

func (v *VPSHUFLW256IDENTITY) Run() {
	vals := [16]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [16]uint16{}

	vpshuflw256Identity(&vals, &ret)

	log.Printf("VPSHUFLW256IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW256IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
