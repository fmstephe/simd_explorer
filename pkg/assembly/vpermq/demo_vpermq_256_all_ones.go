package vpermq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermq_256_all_ones.s
var assemblyVpermq256All_ones string

//go:embed stub_vpermq_256_all_ones.go
var stubVpermq256All_ones string

type VPERMQ256ALL_ONES struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMQ256ALL_ONES() *VPERMQ256ALL_ONES {
	return &VPERMQ256ALL_ONES{
		vals: number.NewNamedUintParameter("vals", 256, 64, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERMQ256ALL_ONES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMQ256ALL_ONES) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMQ256ALL_ONES) Name() string {
	return "VPERMQ (256 bit) all_ones"
}

func (v *VPERMQ256ALL_ONES) Description() string {
	return "Permute 4 u64 elements using imm8=0x55 (all ones)."
}

func (v *VPERMQ256ALL_ONES) Stub() string {
	return stubVpermq256All_ones
}

func (v *VPERMQ256ALL_ONES) Assembly() string {
	return assemblyVpermq256All_ones
}

func (v *VPERMQ256ALL_ONES) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))

	ret := [4]uint64{}

	vpermq256All_ones(&vals, &ret)

	log.Printf("VPERMQ256ALL_ONES vals %v ret %v", vals, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMQ256ALL_ONES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
