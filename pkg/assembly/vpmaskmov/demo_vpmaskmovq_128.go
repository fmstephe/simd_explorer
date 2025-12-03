package vpmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaskmovq_128.s
var assemblyVpmaskmovq128 string

//go:embed stub_vpmaskmovq_128.go
var stubVpmaskmovq128 string

type VPMASKMOVQ128 struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVPMASKMOVQ128() *VPMASKMOVQ128 {
	return &VPMASKMOVQ128{
		vals: number.NewNamedUintParameter("vals", 128, 64, 10),
		mask: number.NewNamedUintParameter("mask", 128, 64, 16),
		ret:  number.NewNamedUintParameter("ret", 128, 64, 10),
	}
}

func (v *VPMASKMOVQ128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VPMASKMOVQ128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMASKMOVQ128) Name() string {
	return "VPMASKMOVQ (128 bit) "
}

func (v *VPMASKMOVQ128) Description() string {
	return "Masked store of packed 64-bit integers: store lanes where mask sign-bit is set."
}

func (v *VPMASKMOVQ128) Stub() string {
	return stubVpmaskmovq128
}

func (v *VPMASKMOVQ128) Assembly() string {
	return assemblyVpmaskmovq128
}

func (v *VPMASKMOVQ128) Run() {
	vals := [2]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	mask := [2]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))

	ret := [2]uint64{}

	vpmaskmovq128(&vals, &mask, &ret)

	log.Printf("VPMASKMOVQ128 vals %v mask %v ret %v", vals, mask, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMASKMOVQ128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
