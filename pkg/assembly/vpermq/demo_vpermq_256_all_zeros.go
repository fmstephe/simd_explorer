package vpermq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermq_256_all_zeros.s
var assemblyVpermq256All_zeros string

//go:embed stub_vpermq_256_all_zeros.go
var stubVpermq256All_zeros string

type VPERMQ256ALL_ZEROS struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMQ256ALL_ZEROS() *VPERMQ256ALL_ZEROS {
	return &VPERMQ256ALL_ZEROS{
		vals: number.NewNamedUintParameter("vals", 256, 64, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERMQ256ALL_ZEROS) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMQ256ALL_ZEROS) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMQ256ALL_ZEROS) Name() string {
	return "VPERMQ (256 bit) all_zeros"
}

func (v *VPERMQ256ALL_ZEROS) Description() string {
	return "Permute 4 u64 elements using imm8=0x00 (all zeros)."
}

func (v *VPERMQ256ALL_ZEROS) Stub() string {
	return stubVpermq256All_zeros
}

func (v *VPERMQ256ALL_ZEROS) Assembly() string {
	return assemblyVpermq256All_zeros
}

func (v *VPERMQ256ALL_ZEROS) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))

	ret := [4]uint64{}

	vpermq256All_zeros(&vals, &ret)

	log.Printf("VPERMQ256ALL_ZEROS vals %v ret %v", vals, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMQ256ALL_ZEROS) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
