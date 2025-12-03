package vpmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaskmovd_256.s
var assemblyVpmaskmovd256 string

//go:embed stub_vpmaskmovd_256.go
var stubVpmaskmovd256 string

type VPMASKMOVD256 struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVPMASKMOVD256() *VPMASKMOVD256 {
	return &VPMASKMOVD256{
		vals: number.NewNamedUintParameter("vals", 256, 32, 10),
		mask: number.NewNamedUintParameter("mask", 256, 32, 16),
		ret:  number.NewNamedUintParameter("ret", 256, 32, 10),
	}
}

func (v *VPMASKMOVD256) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VPMASKMOVD256) Output() *number.Parameter {
	return v.ret
}

func (v *VPMASKMOVD256) Name() string {
	return "VPMASKMOVD (256 bit) "
}

func (v *VPMASKMOVD256) Description() string {
	return "Masked store of packed 32-bit integers: store lanes where mask sign-bit is set."
}

func (v *VPMASKMOVD256) Stub() string {
	return stubVpmaskmovd256
}

func (v *VPMASKMOVD256) Assembly() string {
	return assemblyVpmaskmovd256
}

func (v *VPMASKMOVD256) Run() {
	vals := [8]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	mask := [8]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))

	ret := [8]uint32{}

	vpmaskmovd256(&vals, &mask, &ret)

	log.Printf("VPMASKMOVD256 vals %v mask %v ret %v", vals, mask, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMASKMOVD256) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
