package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_128_all_one.s
var assemblyVpermilpd128All_one string

//go:embed stub_vpermilpd_128_all_one.go
var stubVpermilpd128All_one string

type VPERMILPD128ALL_ONE struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPD128ALL_ONE() *VPERMILPD128ALL_ONE {
	return &VPERMILPD128ALL_ONE{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VPERMILPD128ALL_ONE) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPD128ALL_ONE) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPD128ALL_ONE) Name() string {
	return "VPERMILPD (128 bit) all_one"
}

func (v *VPERMILPD128ALL_ONE) Description() string {
	return "Permute with imm8=0x55: broadcast a1 to all lanes."
}

func (v *VPERMILPD128ALL_ONE) Stub() string {
	return stubVpermilpd128All_one
}

func (v *VPERMILPD128ALL_ONE) Assembly() string {
	return assemblyVpermilpd128All_one
}

func (v *VPERMILPD128ALL_ONE) Run() {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vpermilpd128All_one(&vals, &ret)

	log.Printf("VPERMILPD128ALL_ONE vals %v ret %v", vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)

}

func (v *VPERMILPD128ALL_ONE) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
