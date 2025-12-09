package pshuflw

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpshuflw_128_identity.s
var assemblyVpshuflw128Identity string

//go:embed stub_vpshuflw_128_identity.go
var stubVpshuflw128Identity string

type VPSHUFLW128IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPSHUFLW128IDENTITY() *VPSHUFLW128IDENTITY {
	return &VPSHUFLW128IDENTITY{
		vals: number.NewNamedUintParameter("vals", 128, 16, 10),
		ret:  number.NewNamedUintParameter("ret", 128, 16, 10),
	}
}

func (v *VPSHUFLW128IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPSHUFLW128IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPSHUFLW128IDENTITY) Name() string {
	return "VPSHUFLW (128 bit) identity"
}

func (v *VPSHUFLW128IDENTITY) Description() string {
	return "Shuffle low words in 128-bit lane using imm8=0xE4 (identity); high words unchanged."
}

func (v *VPSHUFLW128IDENTITY) Stub() string {
	return stubVpshuflw128Identity
}

func (v *VPSHUFLW128IDENTITY) Assembly() string {
	return assemblyVpshuflw128Identity
}

func (v *VPSHUFLW128IDENTITY) Run() {
	vals := [8]uint16{}
	copy(vals[:], number.ToUint16Slice(v.vals.FlatData()))

	ret := [8]uint16{}

	vpshuflw128Identity(&vals, &ret)

	log.Printf("VPSHUFLW128IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint16SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPSHUFLW128IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
