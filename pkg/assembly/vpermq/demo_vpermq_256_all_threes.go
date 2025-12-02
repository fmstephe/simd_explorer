package vpermq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermq_256_all_threes.s
var assemblyVpermq256All_threes string

//go:embed stub_vpermq_256_all_threes.go
var stubVpermq256All_threes string

type VPERMQ256ALL_THREES struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMQ256ALL_THREES() *VPERMQ256ALL_THREES {
	return &VPERMQ256ALL_THREES{
		vals: number.NewNamedUintParameter("vals", 256, 64, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERMQ256ALL_THREES) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMQ256ALL_THREES) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMQ256ALL_THREES) Name() string {
	return "VPERMQ (256 bit) all_threes"
}

func (v *VPERMQ256ALL_THREES) Description() string {
	return "Permute 4 u64 elements using imm8=0xFF (all threes)."
}

func (v *VPERMQ256ALL_THREES) Stub() string {
	return stubVpermq256All_threes
}

func (v *VPERMQ256ALL_THREES) Assembly() string {
	return assemblyVpermq256All_threes
}

func (v *VPERMQ256ALL_THREES) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))

	ret := [4]uint64{}

	vpermq256All_threes(&vals, &ret)

	log.Printf("VPERMQ256ALL_THREES vals %v ret %v", vals, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMQ256ALL_THREES) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
