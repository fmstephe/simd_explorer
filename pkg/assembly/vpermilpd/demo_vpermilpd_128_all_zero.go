package vpermilpd

import (
	_ "embed"
	"log"

	"github.com/fmstephe/simd_explorer/pkg/assembly/asmutil"
	"github.com/fmstephe/simd_explorer/pkg/ui/number"
)

//go:embed asm_vpermilpd_128_all_zero.s
var assemblyVpermilpd128All_zero string

//go:embed stub_vpermilpd_128_all_zero.go
var stubVpermilpd128All_zero string

type VPERMILPD128ALL_ZERO struct {
	vals *number.Parameter
	ret  *number.Parameter
}

func NewVPERMILPD128ALL_ZERO() *VPERMILPD128ALL_ZERO {
	return &VPERMILPD128ALL_ZERO{
		vals: number.NewNamedFloatParameter("vals", 128, 64),
		ret:  number.NewNamedFloatParameter("ret", 128, 64),
	}
}

func (v *VPERMILPD128ALL_ZERO) Inputs() []*number.Parameter {
	return []*number.Parameter{
		v.vals,
	}
}

func (v *VPERMILPD128ALL_ZERO) Output() *number.Parameter {
	return v.ret
}

func (v *VPERMILPD128ALL_ZERO) Name() string {
	return "VPERMILPD (128 bit) all_zero"
}

func (v *VPERMILPD128ALL_ZERO) Description() string {
	return "Permute with imm8=0x00: broadcast a0 to all lanes."
}

func (v *VPERMILPD128ALL_ZERO) Stub() string {
	return stubVpermilpd128All_zero
}

func (v *VPERMILPD128ALL_ZERO) Assembly() string {
	return assemblyVpermilpd128All_zero
}

func (v *VPERMILPD128ALL_ZERO) Run() (output []byte) {
	vals := [2]float64{}
	copy(vals[:], number.ToFloat64Slice(v.vals.FlatData()))

	ret := [2]float64{}

	vpermilpd128All_zero(&vals, &ret)

	log.Printf("VPERMILPD128ALL_ZERO vals %v ret %v", vals, ret)

	out := number.Float64SliceToBytes(ret[:])
	v.ret.SetData(out)
	return out
}

func (v *VPERMILPD128ALL_ZERO) Supported() bool {
	return asmutil.IsSupported(v.Assembly())
}
