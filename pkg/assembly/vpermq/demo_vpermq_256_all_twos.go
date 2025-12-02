package vpermq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermq_256_all_twos.s
var assemblyVpermq256All_twos string

//go:embed stub_vpermq_256_all_twos.go
var stubVpermq256All_twos string

type VPERMQ256ALL_TWOS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMQ256ALL_TWOS() *VPERMQ256ALL_TWOS {
	return &VPERMQ256ALL_TWOS{
		vals: number.NewNamedUintParameter("vals", 256, 64, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERMQ256ALL_TWOS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMQ256ALL_TWOS) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMQ256ALL_TWOS) Name() string {
	return "VPERMQ (256 bit) all_twos"
}

func (v *VPERMQ256ALL_TWOS) Description() string {
	return "Permute 4 u64 elements using imm8=0xAA (all twos)."
}

func (v *VPERMQ256ALL_TWOS) Stub() string {
	return stubVpermq256All_twos
}

func (v *VPERMQ256ALL_TWOS) Assembly() string {
	return assemblyVpermq256All_twos
}

func (v *VPERMQ256ALL_TWOS) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))

	ret := [4]uint64{}

	vpermq256All_twos(&vals, &ret)

	log.Printf("VPERMQ256ALL_TWOS vals %v ret %v", vals, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMQ256ALL_TWOS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
