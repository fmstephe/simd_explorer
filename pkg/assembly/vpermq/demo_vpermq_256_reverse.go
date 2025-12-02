package vpermq

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermq_256_reverse.s
var assemblyVpermq256Reverse string

//go:embed stub_vpermq_256_reverse.go
var stubVpermq256Reverse string

type VPERMQ256REVERSE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMQ256REVERSE() *VPERMQ256REVERSE {
	return &VPERMQ256REVERSE{
		vals: number.NewNamedUintParameter("vals", 256, 64, 10),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 10),
	}
}

func (v *VPERMQ256REVERSE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMQ256REVERSE) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMQ256REVERSE) Name() string {
	return "VPERMQ (256 bit) reverse"
}

func (v *VPERMQ256REVERSE) Description() string {
	return "Permute 4 u64 elements using imm8=0x1B (reverse order)."
}

func (v *VPERMQ256REVERSE) Stub() string {
	return stubVpermq256Reverse
}

func (v *VPERMQ256REVERSE) Assembly() string {
	return assemblyVpermq256Reverse
}

func (v *VPERMQ256REVERSE) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))

	ret := [4]uint64{}

	vpermq256Reverse(&vals, &ret)

	log.Printf("VPERMQ256REVERSE vals %v ret %v", vals, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPERMQ256REVERSE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
