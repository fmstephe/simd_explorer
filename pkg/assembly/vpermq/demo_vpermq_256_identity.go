package vpermq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermq_256_identity.s
var assemblyVpermq256Identity string

//go:embed stub_vpermq_256_identity.go
var stubVpermq256Identity string

type VPERMQ256IDENTITY struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMQ256IDENTITY() *VPERMQ256IDENTITY {
	return &VPERMQ256IDENTITY{
		vals: number.NewNamedUintParameter("vals", 256, 64, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERMQ256IDENTITY) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMQ256IDENTITY) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMQ256IDENTITY) Name() string {
	return "VPERMQ (256 bit) identity"
}

func (v *VPERMQ256IDENTITY) Description() string {
	return "Permute 4 u64 elements using imm8=0xE4 (identity)."
}

func (v *VPERMQ256IDENTITY) Stub() string {
	return stubVpermq256Identity
}

func (v *VPERMQ256IDENTITY) Assembly() string {
	return assemblyVpermq256Identity
}

func (v *VPERMQ256IDENTITY) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))

	ret := [4]uint64{}

	vpermq256Identity(&vals, &ret)

	log.Printf("VPERMQ256IDENTITY vals %v ret %v", vals, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMQ256IDENTITY) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
