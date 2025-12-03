package vpmaskmov

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpmaskmovd_128.s
var assemblyVpmaskmovd128 string

//go:embed stub_vpmaskmovd_128.go
var stubVpmaskmovd128 string

type VPMASKMOVD128 struct {
	vals *number.Parameter
	mask *number.Parameter
	ret  *number.Parameter
}

func NewVPMASKMOVD128() *VPMASKMOVD128 {
	return &VPMASKMOVD128{
		vals: number.NewNamedUintParameter("vals", 128, 32, 10),
		mask: number.NewNamedUintParameter("mask", 128, 32, 16),
		ret:  number.NewNamedUintParameter("ret", 128, 32, 10),
	}
}

func (v *VPMASKMOVD128) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
		v.mask,
	}
}

func (v *VPMASKMOVD128) Output() *number.Parameter {
	return v.ret
}

func (v *VPMASKMOVD128) Name() string {
	return "VPMASKMOVD (128 bit) "
}

func (v *VPMASKMOVD128) Description() string {
	return "Masked store of packed 32-bit integers: store lanes where mask sign-bit is set."
}

func (v *VPMASKMOVD128) Stub() string {
	return stubVpmaskmovd128
}

func (v *VPMASKMOVD128) Assembly() string {
	return assemblyVpmaskmovd128
}

func (v *VPMASKMOVD128) Run() {
	vals := [4]uint32{}
	copy(vals[:], number.ToUint32Slice(v.vals.FlatData()))
	mask := [4]uint32{}
	copy(mask[:], number.ToUint32Slice(v.mask.FlatData()))

	ret := [4]uint32{}

	vpmaskmovd128(&vals, &mask, &ret)

	log.Printf("VPMASKMOVD128 vals %v mask %v ret %v", vals, mask, ret)

	out := number.Uint32SliceToBytes(ret[:])
	v.ret.SetData(out)
}

func (v *VPMASKMOVD128) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
