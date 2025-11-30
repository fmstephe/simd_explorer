package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_256_all_one.s
var assemblyVpermilpd256All_one string

//go:embed stub_vpermilpd_256_all_one.go
var stubVpermilpd256All_one string

type VPERMILPD256ALL_ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPD256ALL_ONE() *VPERMILPD256ALL_ONE {
	return &VPERMILPD256ALL_ONE{
		vals: number.NewNamedFloatParameter("vals", 256, 64),
		ret:  number.NewNamedFloatParameter("ret", 256, 64),
	}
}

func (v *VPERMILPD256ALL_ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPD256ALL_ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPD256ALL_ONE) Name() string {
	return "VPERMILPD (256 bit) all_one"
}

func (v *VPERMILPD256ALL_ONE) Description() string {
	return "Permute with imm8=0x55 per 128-bit lane: broadcast lane1 element."
}

func (v *VPERMILPD256ALL_ONE) Stub() string {
	return stubVpermilpd256All_one
}

func (v *VPERMILPD256ALL_ONE) Assembly() string {
	return assemblyVpermilpd256All_one
}

func (v *VPERMILPD256ALL_ONE) Run() (output []byte) {
	vals := [4]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [4]float64{}

	vpermilpd256All_one(&vals, &ret)

	log.Printf("VPERMILPD256ALL_ONE vals %v ret %v", vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPD256ALL_ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
