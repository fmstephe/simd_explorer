package vpmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaskmovq_256.s
var assemblyVpmaskmovq256 string

//go:embed stub_vpmaskmovq_256.go
var stubVpmaskmovq256 string

type VPMASKMOVQ256 struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVPMASKMOVQ256() *VPMASKMOVQ256 {
	return &VPMASKMOVQ256{
		vals: number.NewNamedUintParameter("vals", 256, 64, 16),
		mask: number.NewNamedUintParameter("mask", 256, 64, 16),
		ret:  number.NewNamedUintParameter("ret", 256, 64, 16),
	}
}

func (v *VPMASKMOVQ256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VPMASKMOVQ256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMASKMOVQ256) Name() string {
	return "VPMASKMOVQ (256 bit) "
}

func (v *VPMASKMOVQ256) Description() string {
	return "Masked store of packed 64-bit integers: store lanes where mask sign-bit is set."
}

func (v *VPMASKMOVQ256) Stub() string {
	return stubVpmaskmovq256
}

func (v *VPMASKMOVQ256) Assembly() string {
	return assemblyVpmaskmovq256
}

func (v *VPMASKMOVQ256) Run() {
	vals := [4]uint64{}
	copy(vals[:], number.ToUint64Slice(v.vals.FlatData()))
	mask := [4]uint64{}
	copy(mask[:], number.ToUint64Slice(v.mask.FlatData()))

	ret := [4]uint64{}

	vpmaskmovq256(&vals, &mask, &ret)

	log.Printf("VPMASKMOVQ256 vals %v mask %v ret %v", vals, mask, ret)

	out := number.Uint64SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMASKMOVQ256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
